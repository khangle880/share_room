package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		log.Error().Msg("gin.Context has wrong type")
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}