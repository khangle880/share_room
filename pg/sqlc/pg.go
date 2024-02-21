package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"

	"github.com/khangle880/share_room/utils"
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
		resProfile, err := q.CreateProfile(ctx, profileParams)
		if err != nil {
			return err
		}
		utils.Log.Info().Msg(resProfile.ID.String())
		err = q.CreateUserProfile(ctx, CreateUserProfileParams{
			UserID:    resUser.ID,
			ProfileID: resProfile.ID,
		})
		if err != nil {
			return err
		}
		a := User(resUser)
		user = &a
		return nil
	})
	return user, err
}

type BudgetMemberInput struct {
	ID   uuid.UUID  `json:"id"`
	Role BudgetRole `json:"role"`
}

func (r *RepoSvc) CreateBudget(ctx context.Context, budgetParams CreateBudgetParams, memberIDs []BudgetMemberInput) (*Budget, error) {
	budget := new(Budget)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateBudget(ctx, budgetParams)
		if err != nil {
			return err
		}
		for _, member := range memberIDs {
			if err = q.SetBudgetMember(ctx, SetBudgetMemberParams{
				UserID:   member.ID,
				BudgetID: res.ID,
				Role:     member.Role,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return budget, err
}

func (r *RepoSvc) UpdateBudget(ctx context.Context, budgetParams UpdateBudgetParams, memberIDs []BudgetMemberInput) (*Budget, error) {
	budget := new(Budget)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.UpdateBudget(ctx, budgetParams)
		if err != nil {
			return err
		}
		if memberIDs == nil {
			return nil
		}
		if err = q.UnsetBudgetMembers(ctx, res.ID); err != nil {
			return err
		}
		for _, member := range memberIDs {
			if err = q.SetBudgetMember(ctx, SetBudgetMemberParams{
				UserID:   member.ID,
				BudgetID: res.ID,
				Role:     member.Role,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return budget, err
}

func (r *RepoSvc) CreateTransaction(ctx context.Context, transactionParams CreateTransactionParams, creatorIDs []uuid.UUID, partnerIDs []uuid.UUID) (*Transaction, error) {
	tran := new(Transaction)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateTransaction(ctx, transactionParams)
		if err != nil {
			return err
		}
		// creatorIDs
		for _, creator := range creatorIDs {
			if err = q.SetTranMember(ctx, SetTranMemberParams{
				UserID:        creator,
				TransactionID: res.ID,
				Role:          TransRoleCreator,
			}); err != nil {
				return err
			}
		}
		// partnerIDs
		for _, partner := range partnerIDs {
			if err = q.SetTranMember(ctx, SetTranMemberParams{
				UserID:        partner,
				TransactionID: res.ID,
				Role:          TransRolePartner,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return tran, err
}

func (r *RepoSvc) UpdateTransaction(ctx context.Context, tranParams UpdateTransactionParams, creatorIDs []uuid.UUID, partnerIDs []uuid.UUID) (*Transaction, error) {
	tran := new(Transaction)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.UpdateTransaction(ctx, tranParams)
		if err != nil {
			return err
		}
		if creatorIDs == nil || partnerIDs == nil {
			return nil
		}
		if err = q.UnsetTranMembers(ctx, res.ID); err != nil {
			return err
		}
		// creatorIDs
		for _, creator := range creatorIDs {
			if err = q.SetTranMember(ctx, SetTranMemberParams{
				UserID:        creator,
				TransactionID: res.ID,
				Role:          TransRoleCreator,
			}); err != nil {
				return err
			}
		}
		// partnerIDs
		for _, partner := range partnerIDs {
			if err = q.SetTranMember(ctx, SetTranMemberParams{
				UserID:        partner,
				TransactionID: res.ID,
				Role:          TransRolePartner,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return tran, err
}

func (r *RepoSvc) CreateRoom(ctx context.Context, params CreateRoomParams, adminID uuid.UUID, memberIDs []uuid.UUID) (*Room, error) {
	room := new(Room)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.CreateRoom(ctx, params)
		if err != nil {
			return err
		}
		if err = q.SetRoomMember(ctx, SetRoomMemberParams{
			UserID: adminID,
			RoomID: res.ID,
			Role:   RoomRoleAdmin,
		}); err != nil {
			return err
		}
		for _, member := range memberIDs {
			if err = q.SetRoomMember(ctx, SetRoomMemberParams{
				UserID: member,
				RoomID: res.ID,
				Role:   RoomRoleMember,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return room, err
}

func (r *RepoSvc) UpdateRoom(ctx context.Context, params UpdateRoomParams, adminID *uuid.UUID, memberIDs []uuid.UUID) (*Room, error) {
	room := new(Room)
	err := r.withTx(ctx, func(q *Queries) error {
		res, err := q.UpdateRoom(ctx, params)
		if err != nil {
			return err
		}
		if memberIDs == nil {
			return nil
		}
		if err = q.UnsetRoomMembers(ctx, res.ID); err != nil {
			return err
		}
		if adminID != nil {
			if err = q.SetRoomMember(ctx, SetRoomMemberParams{
				UserID: *adminID,
				RoomID: res.ID,
				Role:   RoomRoleAdmin,
			}); err != nil {
				return err
			}
		}
		for _, member := range memberIDs {
			if err = q.SetRoomMember(ctx, SetRoomMemberParams{
				UserID: member,
				RoomID: res.ID,
				Role:   RoomRoleMember,
			}); err != nil {
				return err
			}
		}
		return nil
	})
	return room, err
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

func UUIDPtrToNullUUID(u *uuid.UUID) uuid.NullUUID {
	if u != nil {
		return uuid.NullUUID{UUID: *u, Valid: true}
	}
	return uuid.NullUUID{}
}

type Decimal decimal.Decimal
type NullDecimal decimal.NullDecimal

func (d Decimal) NumericValue() (pgtype.Numeric, error) {
	dd := decimal.Decimal(d)
	return pgtype.Numeric{Int: dd.Coefficient(), Exp: dd.Exponent(), Valid: true}, nil
}
func (d NullDecimal) NumericValue() (pgtype.Numeric, error) {
	if !d.Valid {
		return pgtype.Numeric{}, nil
	}

	dd := decimal.Decimal(d.Decimal)
	return pgtype.Numeric{Int: dd.Coefficient(), Exp: dd.Exponent(), Valid: true}, nil
}

func NumbericToFloat64(obj pgtype.Numeric) (float64, error) {
	f8, err := obj.Float64Value()
	if err != nil {
		utils.Log.Err(err)
		return 0, errors.New("amount invalid")
	}
	return f8.Float64, nil
}

func Float64ToNumberic(obj *float64) (pgtype.Numeric, error) {
	if obj != nil {
		n := decimal.NewFromFloat(*obj)
		return Decimal(n).NumericValue()
	}
	return pgtype.Numeric{}, nil
}
