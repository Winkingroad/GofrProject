package main

import (
    "context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Winkingroad/GofrProject/stores"
	"github.com/Winkingroad/GofrProject/handlers"
	"github.com/Winkingroad/GofrProject/middleware"
	"github.com/Winkingroad/GofrProject/stores/cars"
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

	store := cars.New() 

	h := handlers.New(store) 

	stores.SetMongoClient(client)
	handlers.SetMongoClient(client)

	app.POST("/login", handlers.LoginHandler)
    app.POST("/signup", handlers.RegisterHandler)
    app.GET("/cars", middleware.JWTAuth(h.GetCars))
    app.GET("/cars/{carno}", middleware.JWTAuth(func(ctx *gofr.Context) (interface{}, error) {return h.Get(ctx, ctx.PathParam("carno"))}))
    app.POST("/cars", middleware.JWTAuth(h.Create))
    app.PUT("/cars/{carno}", middleware.JWTAuth(func(ctx *gofr.Context) (interface{}, error) {return h.Update(ctx, ctx.PathParam("carno"))}))
    app.DELETE("/cars/{carno}", middleware.JWTAuth(func(ctx *gofr.Context) (interface{}, error) {return h.Delete(ctx, ctx.PathParam("carno"))}))
   

	app.Server.HTTP.Port = 9000

	app.Start()
}
