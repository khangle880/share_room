package model

import (
	"time"

	"github.com/google/uuid"
)

// Transaction type definition
type Transaction struct {
	ID          uuid.UUID   `json:"id"`
	CategoryId  uuid.UUID   `json:"categoryId"`
	Description *string     `json:"description,omitempty"`
	Time        time.Time   `json:"time"`
	BudgetId    uuid.UUID   `json:"budgetId"`
	CreatorIds  []uuid.UUID `json:"creatorIds"`
	PartnerIds  []uuid.UUID `json:"partnerIds,omitempty"`
	EventId     *uuid.UUID  `json:"eventId,omitempty"`
	Images      []string    `json:"images,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   *time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"-" pg:",soft_delete"`
}
