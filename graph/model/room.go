package model

import (
	"time"

	"github.com/google/uuid"
)

// Room type definition
type Room struct {
	ID         uuid.UUID   `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	DeletedAt  *time.Time  `json:"-" pg:",soft_delete"`
	Name       string      `json:"name"`
	Address    *string     `json:"address,omitempty"`
	AdminId    uuid.UUID   `json:"adminId"`
	MemberIds  []uuid.UUID `json:"memberIds"`
	Avatar     *string     `json:"avatar,omitempty"`
	Background *string     `json:"background,omitempty"`
}
