package middleware

import (
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
)

var secretKey = []byte("secret")

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTAuth(next gofr.Handler) gofr.Handler {
	return func(ctx *gofr.Context) (interface{}, error) {
		authHeader := ctx.Header("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason:     "Invalid credentials",
			}
		}

		return next(ctx)
	}
}
