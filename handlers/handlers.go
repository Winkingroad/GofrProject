package handlers

import (
	"fmt"
    "encoding/json"
	"ZopSmartproject/models"
	"ZopSmartproject/stores"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
)

type handler struct{

	store stores.Car
}

func New (store stores.Car) handler{
	return handler{
		store: store,
	}

  
}
type CarAlreadyExistsError struct {
    ResponseBody  *errors.Response
}

func (e CarAlreadyExistsError) Error() interface{} {
    return e.ResponseBody
}


func (h handler) GetCars(ctx *gofr.Context) (interface{}, error) {
    resp, err := h.store.GetAllCars(ctx)
    if err != nil {
        return nil, err
    }
    return resp, nil
}


func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("carno")

	resp, err := h.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		return nil, errors.EntityNotFound{Entity: "car", ID: id}
	}

	return resp, nil
}


func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var newCar models.Cars

	err := ctx.Bind(&newCar)
	if err != nil {
		return "", errors.InvalidParam{Param: []string{"body"}}
	}

	// Perform specific field validations
	if newCar.Brand == "" {
		return "", &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Brand field is required",
		}
	}
	if newCar.Model == "" {
		return "", &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Model field is required",
		}
	}
	if newCar.CarNo == "" {
		return "", &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "CarNo field is required",
		}
	}
	if newCar.Year == 0 {
		return "", &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Year field is required and must be greater than zero",
		}
	}
	if newCar.Price == 0 {
		return "", &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Price field is required and must be greater than zero",
		}
	}

	// Check if the car with the given carno already exists in the database
	existingCar, err := h.store.Get(ctx, newCar.CarNo)
	if err != nil {
		return "", err
	}

	if len(existingCar) > 0 {
		return "", &errors.Response{
			StatusCode: 200,
			Code:       "200",
			Reason:     "Car already exists",
		}
	}

	// Proceed to create the car
	err = h.store.Create(ctx, newCar)
	if err != nil {
		return "", err
	}

	updatedCar, err := h.store.Get(ctx, newCar.CarNo)
	if err != nil {
		return "", err
	}

	jsonResp, err := json.Marshal(updatedCar)
	if err != nil {
		return "", err
	}

	return string(jsonResp), nil
}



func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("carno")

	resp, oops := h.store.Get(ctx, id)
	if oops != nil {
		return nil, oops
	}

	if len(resp) == 0 {
		return nil, errors.EntityNotFound{Entity: "car", ID: id}
	}
	

	var c models.Cars
	err := ctx.Bind(&c)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
    
	

	

	err = h.store.Update(ctx, id, c)
	if err != nil {
		return nil, err // Return the original error
	}

	updatedCar, err := h.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	jsonResp, err := json.Marshal(updatedCar)
	if err != nil {
		return "", err
	}

	return string(jsonResp), nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("carno")

	resp, oops := h.store.Get(ctx, id)
	if oops != nil {
		return nil, oops
	}

	if len(resp) == 0 {
		return nil, errors.EntityNotFound{Entity: "car", ID: id}
	}

	_, err := h.store.Delete(ctx, id)
	if err != nil {
		return nil, err 
	}

	return fmt.Sprintf("%v car Removed!", id), nil
}
