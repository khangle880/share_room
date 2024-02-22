package graph

import (
	"github.com/khangle880/share_room/dataloaders"
	"github.com/khangle880/share_room/pg/sqlc"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository  *pg.RepoSvc
	DataLoaders dataloaders.Retriever
}
