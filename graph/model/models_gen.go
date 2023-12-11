// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CreateBudgetInput struct {
	Name           string      `json:"name"`
	Description    *string     `json:"description,omitempty"`
	Balance        int         `json:"balance"`
	TransactionIDs []uuid.UUID `json:"transactionIDs,omitempty"`
	IconID         uuid.UUID   `json:"iconID"`
	MemberIDs      []uuid.UUID `json:"memberIDs,omitempty"`
}

type CreateCategoryInput struct {
	Name     string           `json:"name"`
	Type     CategoryTypeEnum `json:"type"`
	IconID   uuid.UUID        `json:"iconID"`
	ParentID *uuid.UUID       `json:"parentID,omitempty"`
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
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	Email     *string       `json:"email,omitempty"`
	Phone     *string       `json:"phone,omitempty"`
	FirstName *string       `json:"firstName,omitempty"`
	LastName  *string       `json:"lastName,omitempty"`
	Role      *UserRoleEnum `json:"role,omitempty"`
	Bio       *string       `json:"bio,omitempty"`
	Avatar    *string       `json:"avatar,omitempty"`
}

// Icon type definition
type Icon struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	Type      *string    `json:"type,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// User type definition
type User struct {
	ID             uuid.UUID    `json:"id"`
	Username       string       `json:"username"`
	HashedPassword string       `json:"hashedPassword"`
	Email          *string      `json:"email,omitempty"`
	Phone          *string      `json:"phone,omitempty"`
	FirstName      *string      `json:"firstName,omitempty"`
	LastName       *string      `json:"lastName,omitempty"`
	Role           UserRoleEnum `json:"role"`
	Bio            *string      `json:"bio,omitempty"`
	Avatar         *string      `json:"avatar,omitempty"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
}

type CategoryTypeEnum string

const (
	CategoryTypeEnumIncome  CategoryTypeEnum = "INCOME"
	CategoryTypeEnumOutcome CategoryTypeEnum = "OUTCOME"
)

var AllCategoryTypeEnum = []CategoryTypeEnum{
	CategoryTypeEnumIncome,
	CategoryTypeEnumOutcome,
}

func (e CategoryTypeEnum) IsValid() bool {
	switch e {
	case CategoryTypeEnumIncome, CategoryTypeEnumOutcome:
		return true
	}
	return false
}

func (e CategoryTypeEnum) String() string {
	return string(e)
}

func (e *CategoryTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CategoryTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CategoryTypeEnum", str)
	}
	return nil
}

func (e CategoryTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRoleEnum string

const (
	UserRoleEnumAdmin    UserRoleEnum = "ADMIN"
	UserRoleEnumRoommate UserRoleEnum = "ROOMMATE"
	UserRoleEnumCaptain  UserRoleEnum = "CAPTAIN"
)

var AllUserRoleEnum = []UserRoleEnum{
	UserRoleEnumAdmin,
	UserRoleEnumRoommate,
	UserRoleEnumCaptain,
}

func (e UserRoleEnum) IsValid() bool {
	switch e {
	case UserRoleEnumAdmin, UserRoleEnumRoommate, UserRoleEnumCaptain:
		return true
	}
	return false
}

func (e UserRoleEnum) String() string {
	return string(e)
}

func (e *UserRoleEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRoleEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRoleEnum", str)
	}
	return nil
}

func (e UserRoleEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}