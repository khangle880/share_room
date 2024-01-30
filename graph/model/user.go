package model

import (
	"time"

	"github.com/google/uuid"
)

// User type definition
type User struct {
	ID             uuid.UUID  `json:"id"`
	Username       string     `json:"username"`
	HashedPassword string     `json:"hashedPassword"`
	Email          string     `json:"email"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"-" pg:",soft_delete"`
}
