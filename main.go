package main

import (
	"ZopSmartproject/handlers"
	"ZopSmartproject/middleware"
	"ZopSmartproject/stores/cars"
	"gofr.dev/pkg/gofr"
)

func main() {
	// create the application object
	app := gofr.New()

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	store := cars.New()
	h := handlers.New(store)

	// specifying the different routes supported by this service
	app.GET("/cars/{id}", middleware.JWTAuth(h.Get))
	app.POST("/cars", middleware.JWTAuth(h.Create))
	app.PUT("/cars/{id}", middleware.JWTAuth(h.Update))
	app.DELETE("/cars/{id}", middleware.JWTAuth(h.Delete))
	app.Server.HTTP.Port = 8097

	app.Start()
}
