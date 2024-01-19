package graph

import (
	"github.com/khangle880/share_room/graph/model"
	"github.com/khangle880/share_room/postgres/query"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	icons          []*model.Icon
	transactions   []*model.Transaction
	events         []*model.Event
	UsersRepo      query.UsersRepo
	CategoriesRepo query.CategoriesRepo
	BudgetsRepo    query.BudgetRepo
}
