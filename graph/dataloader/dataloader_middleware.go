package dataloader

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/khangle880/share_room/graph/model"
)

type loaderKey string

func DataLoaderMiddleware(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userloader := &UserLoader{
			wait:     1 * time.Millisecond,
			maxBatch: 100,
			fetch: func(keys []string) ([]*model.User, []error) {
				var users []*model.User
				err := db.Model(&users).Where("id in (?)", pg.In(keys)).Select()
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
