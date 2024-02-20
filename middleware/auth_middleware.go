package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/khangle880/share_room/pg/sqlc"
	"github.com/khangle880/share_room/utils"
)

type contextKey string

func AuthMiddleware(repo *pg.RepoSvc) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Next()
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]
		validate, err := utils.JwtValidate(context.Background(), auth)
		if err != nil || !validate.Valid {
			utils.Log.Err(err).Msg("invalid token")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
			return
		}

		customClaims, ok := validate.Claims.(*utils.JwtCustomClaim)
		if !ok {
			c.Next()
			return
		}
		user, err := repo.GetUserByID(c.Request.Context(), customClaims.ID)
		if err != nil {
			utils.Log.Err(err).Msg("user not found")
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), contextKey("auth"), &user)
		profile, err := repo.GetProfileByUserID(c.Request.Context(), customClaims.ID)
		if err != nil {
			c.Next()
			return
		}
		ctx = context.WithValue(ctx, contextKey("profile"), &profile)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetProfileFromContext(ctx context.Context) (*pg.Profile, error) {
	profile, ok := ctx.Value(contextKey("profile")).(*pg.Profile)
	if !ok || profile == nil {
		return nil, errors.New("no profile in context")
	}
	return profile, nil
}

func GetUserFromContext(ctx context.Context) (*pg.User, error) {
	user, ok := ctx.Value(contextKey("auth")).(*pg.User)
	if !ok || user == nil {
		return nil, errors.New("no user in context")
	}
	return user, nil
}
