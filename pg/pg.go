package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/khangle880/share_room/utils"
)

type RepoSvc struct {
	*Queries
	conn *pgx.Conn
}

func (r *RepoSvc) withTx(ctx context.Context, txFn func(*Queries) error) error {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
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

func NewRepository(conn *pgx.Conn) *RepoSvc {
	return &RepoSvc{
		Queries: New(conn),
		conn:      conn,
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
func StringPtrToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{}
}
