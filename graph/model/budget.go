package model

import (
	"github.com/google/uuid"
	"time"
)

// Budget type definition
type Budget struct {
	ID             uuid.UUID      `json:"id"`
	Name           string         `json:"name"`
	Description    *string        `json:"description,omitempty"`
	Balance        int            `json:"balance"`
	TransactionIDs []uuid.UUID    `json:"transactionIDs,omitempty"`
	Transactions   []*Transaction `json:"transactions,omitempty"`
	IconID         uuid.UUID      `json:"iconID"`
	Icon           *Icon          `json:"icon"`
	MemberIDs      []uuid.UUID    `json:"memberIDs,omitempty"`
	Members        []*User        `json:"members,omitempty"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
}
