package model

import (
	"time"

	"github.com/google/uuid"
)

// Event type definition
type Event struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IconID      uuid.UUID  `json:"iconID"`
	Icon        *Icon      `json:"icon"`
	Background  *string    `json:"background,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
