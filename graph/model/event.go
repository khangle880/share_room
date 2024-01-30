package model

import (
	"time"

	"github.com/google/uuid"
)

// Event type definition
type Event struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"-" pg:",soft_delete"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IconId      uuid.UUID  `json:"iconId"`
	Background  *string    `json:"background,omitempty"`
}
