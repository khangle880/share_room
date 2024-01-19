package model

import (
	"time"

	"github.com/google/uuid"
)

// Icon type definition
type Icon struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	Type      *string    `json:"type,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}