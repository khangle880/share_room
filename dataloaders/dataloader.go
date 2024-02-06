package dataloaders

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/khangle880/share_room/pg/sqlc"
	"github.com/khangle880/share_room/utils"
)

type contextKey string

const key = contextKey("dataloaders")

type Loaders struct {
	UserLoader    *UserLoader
	ProfileLoader *ProfileLoader
}

func newLoaders(ctx context.Context, repo *pg.RepoSvc) *Loaders {
	return &Loaders{
		UserLoader:    newUserLoader(ctx, repo),
		ProfileLoader: newProfileLoader(ctx, repo),
	}
}

func newUserLoader(ctx context.Context, repo *pg.RepoSvc) *UserLoader {
	return &UserLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.User, []error) {
			// users, err := db.GetUsersByIDs(c.Request.Context(), utils.ConvertList(keys, func(key string) uuid.UUID {
			// 	return uuid.MustParse(key)
			// }))
			users, err := repo.GetUsersByIDs(ctx, keys)
			return utils.ConvertList(users, func(user pg.User) *pg.User {
				return &user
			}), []error{err}
		},
	}
}
func newProfileLoader(ctx context.Context, repo *pg.RepoSvc) *ProfileLoader {
	return &ProfileLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.Profile, []error) {
			// users, err := db.GetUsersByIDs(c.Request.Context(), utils.ConvertList(keys, func(key string) uuid.UUID {
			// 	return uuid.MustParse(key)
			// }))
			profiles, err := repo.GetProfilesByUserIDs(ctx, keys)
			return utils.ConvertList(profiles, func(profile pg.Profile) *pg.Profile {
				return &profile 
			}), []error{err}
		},
	}
}

type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}
