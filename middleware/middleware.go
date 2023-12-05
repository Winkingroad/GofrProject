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

// Define a middleware function for JWT authentication
func JWTAuth(next gofr.Handler) gofr.Handler {
	return func(ctx *gofr.Context) (interface{}, error) {
		// Get the JWT from the request header
		tokenString := ctx.Header("Authorization")

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		// Check if the token is valid
		if err != nil || !token.Valid {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason: "Invalid credentials",
			}
		}

		// Get the claims from the token
		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason: "Invalid credentials",
			}
		}
		

		// Check if the user is authorized to access the endpoint
		if claims.Username != "admin" {
			return nil, &errors.Response{
				StatusCode: 403,
				Reason: "forbidden",
			}
		}

		// Call the next handler function
		return next(ctx)
	}
}
