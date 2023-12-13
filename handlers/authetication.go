package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "ZopSmartproject/stores"

    "github.com/golang-jwt/jwt/v4"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "gofr.dev/pkg/errors"
    "gofr.dev/pkg/gofr"
)

var (
    secretKey    = []byte("secret")
    mongoClient *mongo.Client
)



func SetMongoClient(c *mongo.Client) {
    mongoClient = c
}

type Credentials = stores.Credentials



type User struct {
    Username string
}


func FindUserByUsernameAndPassword(username, password string) (*User, error) {
    collection := mongoClient.Database("cluster21").Collection("users")

    var user User
    err := collection.FindOne(context.Background(), bson.M{"username": username, "password": password}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, &errors.Response{Reason: "User not found"}
        }
        return nil, err
    }

    return &user, nil
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

    validCredentials, err := FindUserByUsernameAndPassword(creds.Username, creds.Password)
    if err != nil {
        return nil, &errors.Response{
            StatusCode: 500,
            Reason:     "Internal server error",
        }
    }

    if validCredentials == nil {
        return nil, &errors.Response{
            StatusCode: 401,
            Reason:     "Invalid credentials",
        }
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &jwt.StandardClaims{
        Subject:   creds.Username,
        ExpiresAt: expirationTime.Unix(),
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

func RegisterHandler(ctx *gofr.Context) (interface{}, error) {
    var newUser Credentials
    err := json.NewDecoder(ctx.Request().Body).Decode(&newUser)
    if err != nil {
        return nil, &errors.Response{
            StatusCode: 400,
            Reason:     "Invalid request body",
        }
    }

    err = stores.InsertUser(newUser)
    if err != nil {
        fmt.Println("Error inserting user:", err) // Log the error
        return nil, &errors.Response{
            StatusCode: 500,
            Reason:     "Failed to register user",
        }
    }

    return "Registration Sucessful", nil
}
