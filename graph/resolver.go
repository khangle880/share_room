package graph

import "github.com/khangle880/share_room/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	icons        []*model.Icon
	users        []*model.User
	budgets      []*model.Budget
	transactions []*model.Transaction
	categories   []*model.Category
	events       []*model.Event
}
