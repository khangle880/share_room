package model

import (
	"time"

	"github.com/google/uuid"
)

// Profile type definition
type Profile struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Role      UserRole   `json:"role"`
	Firstname *string    `json:"firstname,omitempty"`
	Lastname  *string    `json:"lastname,omitempty"`
	Dob       *time.Time `json:"dob,omitempty"`
	Bio       *string    `json:"bio,omitempty"`
	Avatar    *string    `json:"avatar,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
}
