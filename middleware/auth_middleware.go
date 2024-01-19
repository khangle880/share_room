package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khangle880/share_room/graph/model"
	"github.com/khangle880/share_room/postgres/query"
	"github.com/khangle880/share_room/utils"
)

type authUser string

func AuthMiddleware(repo query.UsersRepo) gin.HandlerFunc {
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
			c.JSON(http.StatusForbidden, fmt.Sprintf("Invalid Token: %v", err))
			return
		}

		customClaims, ok := validate.Claims.(*utils.JwtCustomClaim)
		if !ok {
			c.Next()
			return
		}
		user, err := repo.GetUserByID(customClaims.ID)
		if err != nil {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), authUser("auth"), user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetUserFromContext(ctx context.Context) (*model.User, error) {
	user, ok := ctx.Value(authUser("auth")).(*model.User)
	if !ok || user == nil {
		return nil, errors.New("no user in context")
	}
	return user, nil
}
