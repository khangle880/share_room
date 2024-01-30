package model

import (
	"time"

	"github.com/google/uuid"
)

// Budget type definition
type Budget struct {
	ID             uuid.UUID   `json:"id"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAtomitempty"`
	DeletedAt      *time.Time  `json:"-" pg:",soft_delete"`
	Name           string      `json:"name"`
	Description    *string     `json:"description,omitempty"`
	Amount         float64     `json:"amount" pg:",use_zero"`
	IconId         uuid.UUID   `json:"iconId"`
	RoomId         *uuid.UUID  `json:"roomId,omitempty"`
	Period         *PeriodType `json:"period,omitempty"`
	StartDate      time.Time   `json:"startDate"`
	EndDate        *time.Time  `json:"end_date,omitempty"`
	TransactionIds []uuid.UUID `json:"transactionIds,omitempty"`
	MemberIds      []uuid.UUID `json:"memberIds,omitempty"`
}
