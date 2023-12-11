package model

import (
	"github.com/google/uuid"
	"time"
)

// Transaction type definition
type Transaction struct {
	ID          uuid.UUID    `json:"id"`
	CategoryID  uuid.UUID    `json:"categoryID"`
	Category    *Category    `json:"category"`
	Description *string      `json:"description,omitempty"`
	Time        time.Time    `json:"time"`
	BudgetID    uuid.UUID    `json:"budgetID"`
	Budget      *Budget      `json:"budget"`
	CreatorIDs  []uuid.UUID  `json:"creatorIDs"`
	Creators    []*User      `json:"creators"`
	PartnerIDs  []uuid.UUID `json:"partnerIDs,omitempty"`
	Partners    []*User      `json:"partners,omitempty"`
	EventID     *uuid.UUID   `json:"eventID,omitempty"`
	Event       Event        `json:"event,omitempty"`
	Images      []string     `json:"images,omitempty"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty"`
}
