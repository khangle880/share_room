package model

import (
	"time"

	"github.com/google/uuid"
)

// Budget type definition
type Budget struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Balance     int         `json:"balance" pg:",use_zero"`
	IconId      uuid.UUID   `json:"iconId"`
	MemberIds   []uuid.UUID `json:"memberIds,omitempty" pg:",array"`
	RoomId      *uuid.UUID  `json:"roomId,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   *time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
