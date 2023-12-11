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
	IconID    uuid.UUID       `json:"iconID"`
	Icon      *Icon            `json:"icon"`
	ParentID  *uuid.UUID       `json:"parentID,omitempty"`
	Parent    *Category        `json:"parent,omitempty"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt *time.Time       `json:"updatedAt,omitempty"`
}
