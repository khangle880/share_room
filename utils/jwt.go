package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtCustomClaim struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) 

func JwtGenerate(userId uuid.UUID) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtCustomClaim{
		ID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
