// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/khangle880/share_room/pg"
)

// Query Input
type BudgetFilter struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateBudgetInput struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Amount      float64     `json:"Amount"`
	IconID      uuid.UUID   `json:"iconID"`
	RoomID      *uuid.UUID  `json:"roomID,omitempty"`
	MemberIDs   []uuid.UUID `json:"memberIDs,omitempty"`
}

type CreateCategoryInput struct {
	Name     string          `json:"name"`
	Type     pg.CategoryType `json:"type"`
	IconID   uuid.UUID       `json:"iconID"`
	ParentID *uuid.UUID      `json:"parentID,omitempty"`
}

type CreateEventInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IconID      uuid.UUID `json:"iconID"`
	Background  *string   `json:"background,omitempty"`
}

type CreateIconInput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CreateTransInput struct {
	CategoryID  uuid.UUID   `json:"CategoryID"`
	Description *string     `json:"description,omitempty"`
	Time        time.Time   `json:"time"`
	BudgetID    uuid.UUID   `json:"budgetID"`
	CreatorIDs  []uuid.UUID `json:"creatorIDs"`
	PartnerIDs  []uuid.UUID `json:"partnerIDs,omitempty"`
	EventID     *uuid.UUID  `json:"eventID,omitempty"`
	Images      []string    `json:"images,omitempty"`
}

// Input
type CreateUserInput struct {
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Email     *string      `json:"email,omitempty"`
	Phone     *string      `json:"phone,omitempty"`
	Firstname *string      `json:"firstname,omitempty"`
	Lastname  *string      `json:"lastname,omitempty"`
	Role      *pg.UserRole `json:"role,omitempty"`
	Bio       *string      `json:"bio,omitempty"`
	Avatar    *string      `json:"avatar,omitempty"`
}

// Event type definition
type Event struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        *pg.Icon   `json:"icon"`
	Background  *string    `json:"background,omitempty"`
}

type Token struct {
	AccessToken  *string  `json:"accessToken,omitempty"`
	RefreshToken *string  `json:"refreshToken,omitempty"`
	User         *pg.User `json:"user,omitempty"`
}

type UpdateBudgetInput struct {
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Amount      *float64    `json:"Amount,omitempty"`
	IconID      *uuid.UUID  `json:"iconID,omitempty"`
	RoomID      *uuid.UUID  `json:"roomID,omitempty"`
	MemberIDs   []uuid.UUID `json:"memberIDs,omitempty"`
}

type UpdateUserInput struct {
	Username  *string      `json:"username,omitempty"`
	Password  *string      `json:"password,omitempty"`
	Email     string       `json:"email"`
	Phone     *string      `json:"phone,omitempty"`
	Firstname *string      `json:"firstname,omitempty"`
	Lastname  *string      `json:"lastname,omitempty"`
	Role      *pg.UserRole `json:"role,omitempty"`
	Bio       *string      `json:"bio,omitempty"`
	Avatar    *string      `json:"avatar,omitempty"`
}
