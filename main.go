package main

import (
    "context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ZopSmartproject/stores"
	"ZopSmartproject/handlers"
	"ZopSmartproject/middleware"
	"ZopSmartproject/stores/cars"
	"gofr.dev/pkg/gofr"
)

func main() {
	
	app := gofr.New()
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test2:1234567890@cluster21.tfu5ndl.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
if err != nil {
    log.Fatal("Couldn't ping the database:", err)
}

log.Println("Connected to MongoDB!")

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := cars.New(client) 

	h := handlers.New(store) 

	stores.SetMongoClient(client)
	handlers.SetMongoClient(client)

	// specifying the different routes supported by this service
	app.POST("/login", handlers.LoginHandler)
	app.POST("/signup", handlers.RegisterHandler)
	app.GET("/cars", middleware.JWTAuth(h.GetCars))
	app.GET("/cars/{carno}", middleware.JWTAuth(h.Get))
	app.POST("/cars",middleware.JWTAuth(h.Create))
	app.PUT("/cars/{carno}", middleware.JWTAuth(h.Update))
	app.DELETE("/cars/{carno}", middleware.JWTAuth(h.Delete))
	app.Server.HTTP.Port = 8097

	app.Start()
}
