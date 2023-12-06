package handlers

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
	"ZopSmartproject/middleware"
)

var secretKey = []byte("secret")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(ctx *gofr.Context) (interface{}, error) {
	var creds Credentials
	err := json.NewDecoder(ctx.Request().Body).Decode(&creds)
	if err != nil {
		return nil, &errors.Response{
			StatusCode: 400,
			Reason:     "Invalid request body",
		}
	}

	// Example check for valid credentials
	if creds.Username == "admin" && creds.Password == "adminpassword" {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &middleware.CustomClaims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return nil, &errors.Response{
				StatusCode: 401,
				Reason:     "Unable to generate token",
			}
		}

		response := map[string]string{"token": tokenString}
		return response, nil
	}

	return nil, &errors.Response{
		StatusCode: 401,
		Reason:     "Invalid credentials",
	}
}
