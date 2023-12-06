package handlers

import (
	"fmt"
	"log"
	"ZopSmartproject/models"
	"ZopSmartproject/stores"
	"gofr.dev/pkg/gofr"
)

type handler struct{

	store stores.Car
}

func New (store stores.Car) handler{
	return handler{
		store: store,
	}

  
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error){
	id := ctx.PathParam("id")

	resp, err := h.store.Get(ctx,id)
	if err != nil{
		return nil, err
	}
	return resp, nil
}
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var c models.Cars

	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}

	err = h.store.Create(ctx, c)
	if err != nil {
		// Log the error here to understand what went wrong
		log.Println("Error creating car:", err)
		return nil, err
	}

	return "New Car Added!", nil
}
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	var c models.Cars
	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}
	err = h.store.Update(ctx, id, c)
	if err != nil {
		return nil, err
	}
	return "Car Updated!", nil

}


func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	_, err := h.store.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("%v car Removed!", id), nil
}

