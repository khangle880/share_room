// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package pg

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateIcon(ctx context.Context, arg CreateIconParams) (Icon, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) error
	GetProfileByUserID(ctx context.Context, id uuid.UUID) (Profile, error)
	GetProfilesByUserIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]Profile, error)
	GetUserByEmail(ctx context.Context, email *string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUsersByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]User, error)
}

var _ Querier = (*Queries)(nil)