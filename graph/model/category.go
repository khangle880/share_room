package model

import (
	"time"

	"github.com/google/uuid"
)

// Category type definition
type Category struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt *time.Time   `json:"-" pg:",soft_delete"`
	Name      string       `json:"name"`
	Type      CategoryType `json:"type"`
	IconId    uuid.UUID    `json:"iconId"`
	ParentId  *uuid.UUID   `json:"parentId,omitempty"`
}
