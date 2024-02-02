package dataloaders

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khangle880/share_room/pg"
)

func Middleware(repo *pg.RepoSvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		loaders := newLoaders(c.Request.Context(), repo)
		ctx := context.WithValue(c.Request.Context(), key, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}