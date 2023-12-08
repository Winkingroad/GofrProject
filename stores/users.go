package stores

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gofr.dev/pkg/errors"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string
}

var client *mongo.Client // Assuming you have initialized the MongoDB client

func SetMongoClient(c *mongo.Client) {
    client = c
}

func InsertUser(newUser Credentials) error {
	collection := client.Database("cluster21").Collection("users")

	// Check if the username already exists
	existingUser := FindUserByUsername(newUser.Username)
	if existingUser != nil {
		return &errors.Response{
			StatusCode: 400,
			Reason:     "Username already exists",
		}
	}

	_, err := collection.InsertOne(context.Background(), bson.M{"username": newUser.Username, "password": newUser.Password})
	if err != nil {
		return err
	}

	return nil
}

func FindUserByUsername(username string) *User {
	collection := client.Database("cluster21").Collection("users")

	var user User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil // User not found or error occurred
	}

	return &user
}
