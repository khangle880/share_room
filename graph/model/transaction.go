package model

import (
	"time"

	"github.com/google/uuid"
)

// Transaction type definition
type Transaction struct {
	ID          uuid.UUID   `json:"id"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
	CategoryId  uuid.UUID   `json:"categoryId"`
	BudgetId    uuid.UUID   `json:"budgetId"`
	EventId     *uuid.UUID  `json:"eventId,omitempty"`
	Description *string     `json:"description,omitempty"`
	ExcTime     time.Time   `json:"excTime"`
	CreatorIds  []uuid.UUID `json:"creatorIds"`
	PartnerIds  []uuid.UUID `json:"partnerIds,omitempty"`
	Amount      float64    `json:"amount"`
	Images      []string    `json:"images,omitempty"`
}
