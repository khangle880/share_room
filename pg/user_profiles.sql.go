// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: user_profiles.sql

package pg

import (
	"context"

	"github.com/google/uuid"
)

const createUserProfile = `-- name: CreateUserProfile :exec
INSERT INTO user_profiles (user_id, profile_id) 
VALUES ($1, $2)
`

type CreateUserProfileParams struct {
	UserID    uuid.UUID
	ProfileID uuid.UUID
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) error {
	_, err := q.db.ExecContext(ctx, createUserProfile, arg.UserID, arg.ProfileID)
	return err
}
