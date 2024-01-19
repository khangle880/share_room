package model

import (
	"time"

	"github.com/google/uuid"
)

// Room type definition
type Room struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	CaptainId  uuid.UUID  `json:"captainId"`
	Avatar     *string    `json:"avatar,omitempty"`
	Background *string    `json:"background,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
