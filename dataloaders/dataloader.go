package dataloaders

import (
	"context"
	"time"

	"github.com/google/uuid"
	pg "github.com/khangle880/share_room/pg/sqlc"
)

type contextKey string

const key = contextKey("dataloaders")

func mapByKeys[K comparable, V any](data []V, keys []K, keyFunc func(V) K) map[K]*V {
	group := make(map[K]*V, len(keys))
	for _, r := range data {
		key := keyFunc(r)
		group[key] = &r
	}
	return group
}

func groupBy[K comparable, V any](data []V, keys []K, keyFunc func(V) K) map[K][]V {
	group := make(map[K][]V, len(keys))
	for _, r := range data {
		key := keyFunc(r)
		group[key] = append(group[key], r)
	}
	return group
}

func groupByAndConvert[K comparable, V any, C any](data []V, keys []K, keyFunc func(V) K, convert func(V) C) map[K][]C {
	group := make(map[K][]C, len(keys))
	for _, r := range data {
		key := keyFunc(r)
		convertedValue := convert(r)
		group[key] = append(group[key], convertedValue)
	}
	return group
}

func order[K comparable, V any](data map[K]V, keys []K) []V {
	result := make([]V, len(keys))
	for i, key := range keys {
		result[i] = data[key]
	}
	return result
}

type Loaders struct {
	User              *UserLoader
	Icon              *IconLoader
	Profile           *ProfileLoader
	Event             *EventLoader
	TransByBudgetID   *TransactionSliceLoader
	MembersByRoomID   *UserSliceLoader
	MembersByBudgetID *UserSliceLoader
	CreatorsByTranID  *UserSliceLoader
	PartnersByTranID  *UserSliceLoader
}

func newLoaders(ctx context.Context, repo *pg.RepoSvc) *Loaders {
	return &Loaders{
		User:              newUserLoader(ctx, repo),
		Icon:              newIconLoader(ctx, repo),
		Profile:           newProfileLoader(ctx, repo),
		Event:             newEventLoader(ctx, repo),
		TransByBudgetID:   newTransByBudgetIDLoader(ctx, repo),
		MembersByRoomID:   newMembersByRoomIDLoader(ctx, repo),
		MembersByBudgetID: newMembersByBudgetIDLoader(ctx, repo),
		CreatorsByTranID:  newMembersByTranIDAndRoleLoader(ctx, repo, pg.TransRoleCreator),
		PartnersByTranID:  newMembersByTranIDAndRoleLoader(ctx, repo, pg.TransRolePartner),
	}
}

func newUserLoader(ctx context.Context, repo *pg.RepoSvc) *UserLoader {
	return &UserLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.User, []error) {
			// query
			res, err := repo.GetUsersByIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := mapByKeys[uuid.UUID, pg.User](res, keys, func(u pg.User) uuid.UUID { return u.ID })
			ordered := order[uuid.UUID, *pg.User](group, keys)
			return ordered, nil
		},
	}
}
func newProfileLoader(ctx context.Context, repo *pg.RepoSvc) *ProfileLoader {
	return &ProfileLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.Profile, []error) {
			res, err := repo.GetProfilesByUserIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := mapByKeys[uuid.UUID, pg.Profile](res, keys, func(p pg.Profile) uuid.UUID { return p.ID })
			ordered := order[uuid.UUID, *pg.Profile](group, keys)
			return ordered, nil
		},
	}
}
func newIconLoader(ctx context.Context, repo *pg.RepoSvc) *IconLoader {
	return &IconLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.Icon, []error) {
			res, err := repo.GetIconsByIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := mapByKeys[uuid.UUID, pg.Icon](res, keys, func(icon pg.Icon) uuid.UUID { return icon.ID })
			ordered := order[uuid.UUID, *pg.Icon](group, keys)
			return ordered, nil
		},
	}
}

func newEventLoader(ctx context.Context, repo *pg.RepoSvc) *EventLoader {
	return &EventLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([]*pg.Event, []error) {
			res, err := repo.GetEventsByIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := mapByKeys[uuid.UUID, pg.Event](res, keys, func(e pg.Event) uuid.UUID { return e.ID })
			ordered := order[uuid.UUID, *pg.Event](group, keys)
			return ordered, nil
		},
	}
}
func newTransByBudgetIDLoader(ctx context.Context, repo *pg.RepoSvc) *TransactionSliceLoader {
	return &TransactionSliceLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([][]pg.Transaction, []error) {
			res, err := repo.GetTransByBudgetIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := groupBy[uuid.UUID, pg.Transaction](res, keys, func(t pg.Transaction) uuid.UUID { return t.BudgetID.UUID })
			ordered := order[uuid.UUID, []pg.Transaction](group, keys)
			return ordered, nil
		},
	}
}
func newMembersByBudgetIDLoader(ctx context.Context, repo *pg.RepoSvc) *UserSliceLoader {
	return &UserSliceLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([][]pg.User, []error) {
			res, err := repo.GetMembersByBudgetIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := groupByAndConvert[uuid.UUID](res, keys,
				func(t pg.GetMembersByBudgetIDsRow) uuid.UUID { return t.BudgetID }, func(t pg.GetMembersByBudgetIDsRow) pg.User { return t.User })
			ordered := order[uuid.UUID, []pg.User](group, keys)
			return ordered, nil
		},
	}
}
func newMembersByTranIDAndRoleLoader(ctx context.Context, repo *pg.RepoSvc, role pg.TransRole) *UserSliceLoader {
	return &UserSliceLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([][]pg.User, []error) {
			res, err := repo.GetMembersByTranIDs(ctx, pg.GetMembersByTranIDsParams{
				Column1: keys,
				Role:    role,
			})
			if err != nil {
				return nil, []error{err}
			}
			group := groupByAndConvert[uuid.UUID](res, keys,
				func(t pg.GetMembersByTranIDsRow) uuid.UUID { return t.TransactionID }, func(t pg.GetMembersByTranIDsRow) pg.User { return t.User })
			ordered := order[uuid.UUID, []pg.User](group, keys)
			return ordered, nil
		},
	}
}
func newMembersByRoomIDLoader(ctx context.Context, repo *pg.RepoSvc) *UserSliceLoader {
	return &UserSliceLoader{
		wait:     1 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []uuid.UUID) ([][]pg.User, []error) {
			res, err := repo.GetMembersByRoomIDs(ctx, keys)
			if err != nil {
				return nil, []error{err}
			}
			group := groupByAndConvert[uuid.UUID](res, keys,
				func(t pg.GetMembersByRoomIDsRow) uuid.UUID { return t.RoomID }, func(t pg.GetMembersByRoomIDsRow) pg.User { return t.User })
			ordered := order[uuid.UUID, []pg.User](group, keys)
			return ordered, nil
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
