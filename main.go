package main

import (
	"ZopSmartproject/handlers"
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
	app.GET("/cars/{id}", h.Get)
	app.POST("/cars", h.Create)
	app.DELETE("/cars/{id}", h.Delete)
	app.Server.HTTP.Port = 8097

	app.Start()
}