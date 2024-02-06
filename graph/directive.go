package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/khangle880/share_room/middleware"
	"github.com/khangle880/share_room/pg/sqlc"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user, err := middleware.GetUserFromContext(ctx)
	if err != nil || user == nil {
		return nil, errors.New("access denied")
	}
	return next(ctx)
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role pg.UserRole) (interface{}, error) {
	profile, err := middleware.GetProfileFromContext(ctx)
	if err != nil || profile.Role != role {
		return nil, errors.New("access denied")
	}

	return next(ctx)
}
