package dataloader

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/khangle880/share_room/graph/model"
	"github.com/khangle880/share_room/utils"
)

type loaderKey string

func DataLoaderMiddleware(db *database.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userloader := &UserLoader{
			wait:     1 * time.Millisecond,
			maxBatch: 100,
			fetch: func(keys []string) ([]*model.User, []error) {
				users, err := db.GetUsersByIDs(c.Request.Context(), utils.ConvertList(keys, func(key string) uuid.UUID {
					return uuid.MustParse(key)
				}))
				return users, []error{err}
			},
		}

		ctx := context.WithValue(c.Request.Context(), loaderKey("Users"), userloader)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(loaderKey("Users")).(*UserLoader)
}
