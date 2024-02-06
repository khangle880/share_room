package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/khangle880/share_room/utils"
	"github.com/mitchellh/mapstructure"
)

type RepoSvc struct {
	*Queries
	connPool *pgxpool.Pool
}

func (r *RepoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return err
	}
	q := r.Queries.WithTx(tx)
	err = txFn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			err = fmt.Errorf("tx failed: %v, unable to rollback %v", err, rbErr)
		}
	} else {
		err = tx.Commit(ctx)
	}
	return err
}

func NewRepository(connPool *pgxpool.Pool) *RepoSvc {
	return &RepoSvc{
		Queries:  New(connPool),
		connPool: connPool,
	}
}

func (r *RepoSvc) CreateUser(ctx context.Context, userParams CreateUserParams, profileParams CreateProfileParams) (*User, error) {
	user := new(User)
	err := r.withTx(ctx, func(q *Queries) error {
		resUser, err := q.CreateUser(ctx, userParams)
		if err != nil {
			return err
		}
		utils.Log.Info().Msg(resUser.ID.String())
		resProfile, err := r.CreateProfile(ctx, profileParams)
		if err != nil {
			return err
		}
		utils.Log.Info().Msg(resProfile.ID.String())
		err = r.CreateUserProfile(ctx, CreateUserProfileParams{
			UserID:    resUser.ID,
			ProfileID: resProfile.ID,
		})
		if err != nil {
			return err
		}
		user = &resUser
		return nil
	})
	return user, err
}

// Open opens a database specified by the data source name.
// Format: host=foo port=5432 user=bar password=baz dbname=qux sslmode=disable"
func Open(dataSourceName string) (*sql.DB, error) {
	return sql.Open("postgres", dataSourceName)
}

// StringPtrToNullString converts *string to sql.NullString.
// func StringPtrToNullString(s *string) sql.NullString {
// 	if s != nil {
// 		return sql.NullString{String: *s, Valid: true}
// 	}
// 	return sql.NullString{}
// }
// func NullStringToStringPtr(obj sql.NullString) *string {
// 	var s string
// 	if obj.Valid {
// 		s = obj.String
// 		return &s
// 	}
// 	return nil
// }

func NumbericToFloat64(obj pgtype.Numeric) (float64, error) {
	f8, err := obj.Float64Value()
	if err != nil {
		utils.Log.Err(err)
		return 0, errors.New("amount invalid")
	}
	return f8.Float64, nil
}

func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {
			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}