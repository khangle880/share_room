package model

import (
	"time"

	"github.com/google/uuid"
)

// User type definition
type User struct {
	ID             uuid.UUID    `json:"id"`
	Username       string       `json:"username"`
	HashedPassword string       `json:"hashedPassword"`
	Email          string       `json:"email"`
	Phone          *string      `json:"phone,omitempty"`
	Firstname      *string      `json:"firstname,omitempty"`
	Lastname       *string      `json:"lastname,omitempty"`
	Role           UserRoleEnum `json:"role"`
	Bio            *string      `json:"bio,omitempty"`
	Avatar         *string      `json:"avatar,omitempty"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
}
