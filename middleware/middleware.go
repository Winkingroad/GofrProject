package middleware

import (
	

	"github.com/golang-jwt/jwt/v4"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
)

// Define a secret key for signing the JWTs
var secretKey = []byte("my_secret_key")

// Define a custom struct to store the JWT claims
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


func JWTAuth(next gofr.Handler) gofr.Handler {
	return func(ctx *gofr.Context) (interface{}, error) {
		
		tokenString := ctx.Header("Authorization")

		
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		
		if err != nil || !token.Valid {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason: "Invalid credentials",
			}
		}

		
		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason: "Invalid credentials",
			}
		}
		

		
		if claims.Username != "admin" {
			return nil, &errors.Response{
				StatusCode: 403,
				Reason: "forbidden",
			}
		}

		
		return next(ctx)
	}
}
