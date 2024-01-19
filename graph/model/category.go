package model

import (
	"time"

	"github.com/google/uuid"
)

// Category type definition
type Category struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	Type      CategoryTypeEnum `json:"type"`
	IconId    uuid.UUID        `json:"iconId"`
	ParentId  *uuid.UUID       `json:"parentId,omitempty"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt *time.Time       `json:"updatedAt,omitempty"`
	DeletedAt *time.Time       `json:"-" pg:",soft_delete"`
}
