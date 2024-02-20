// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package pg

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateBudget(ctx context.Context, arg CreateBudgetParams) (Budget, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateIcon(ctx context.Context, arg CreateIconParams) (Icon, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) error
	DeleteBudget(ctx context.Context, id uuid.UUID) error
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetBudgetByID(ctx context.Context, id uuid.UUID) (Budget, error)
	GetBudgets(ctx context.Context, arg GetBudgetsParams) ([]Budget, error)
	GetBudgetsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Budget, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategoriesByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Category, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (Category, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (Event, error)
	GetEventsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Event, error)
	GetIconByID(ctx context.Context, id uuid.UUID) (Icon, error)
	GetIconsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Icon, error)
	GetMembersByBudgetIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]GetMembersByBudgetIDsRow, error)
	GetMembersByRoomIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]GetMembersByRoomIDsRow, error)
	GetMembersByTranIDs(ctx context.Context, arg GetMembersByTranIDsParams) ([]GetMembersByTranIDsRow, error)
	GetProfileByUserID(ctx context.Context, id uuid.UUID) (Profile, error)
	GetProfilesByUserIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Profile, error)
	GetRoomByID(ctx context.Context, id uuid.UUID) (Room, error)
	GetRooms(ctx context.Context, arg GetRoomsParams) ([]Room, error)
	GetRoomsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Room, error)
	GetTransByBudgetIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (Transaction, error)
	GetTransactions(ctx context.Context, arg GetTransactionsParams) ([]Transaction, error)
	GetTransactionsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Transaction, error)
	GetUserByEmail(ctx context.Context, email *string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByRoomRole(ctx context.Context, arg GetUserByRoomRoleParams) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUsersByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]User, error)
	SetBudgetMember(ctx context.Context, arg SetBudgetMemberParams) error
	SetRoomMember(ctx context.Context, arg SetRoomMemberParams) error
	SetTranMember(ctx context.Context, arg SetTranMemberParams) error
	UnsetBudgetMembers(ctx context.Context, budgetID uuid.UUID) error
	UnsetRoomMembers(ctx context.Context, roomID uuid.UUID) error
	UnsetTranMembers(ctx context.Context, transactionID uuid.UUID) error
	UpdateBudget(ctx context.Context, arg UpdateBudgetParams) (Budget, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error)
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
