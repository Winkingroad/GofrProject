package cars

import (
	

	"ZopSmartproject/models"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type car struct{}

func New() car {
     return car{}
}


func (c car) GetAllCars(ctx *gofr.Context) ([]models.Cars, error) {
    resp := make([]models.Cars, 0)

	collection := ctx.MongoDB.Collection("cars")

    // Get all cars without filtering
    curr, err := collection.Find(ctx, bson.D{})
    if err != nil {
        return resp, errors.DB{Err: err}
    }
    defer curr.Close(ctx)

    for curr.Next(ctx) {
        var car models.Cars
        err := curr.Decode(&car)
        if err != nil {
            return resp, errors.DB{Err: err}
        }
        resp = append(resp, car)
    }

    return resp, nil
}

func (c car) Get(ctx *gofr.Context, carno string) ([]models.Cars, error) {
    resp := make([]models.Cars, 0)

	collection := ctx.MongoDB.Collection("cars")

    filter := bson.D{{"carno", carno}}

    curr, err := collection.Find(ctx, filter)
    if err != nil {
        return resp, errors.DB{Err: err}
    }
    defer curr.Close(ctx)

    for curr.Next(ctx) {
        var car models.Cars
        err := curr.Decode(&car)
        if err != nil {
            return resp, errors.DB{Err: err}
        }
        resp = append(resp, car)
    }

    return resp, nil
}


func (c car) Create(ctx *gofr.Context, model models.Cars) error {
    collection := ctx.MongoDB.Collection("cars")

    // Check if the carno already exists in the database
    filter := bson.D{{"carno", model.CarNo}}
    existingCar := collection.FindOne(ctx, filter)
    if existingCar.Err() == nil {
       
        return errors.EntityAlreadyExists{}
    }
    _, err := collection.InsertOne(ctx, model)
    if err != nil {
        return errors.DB{Err: err}
    }

    return nil
}


func (c car) Update(ctx *gofr.Context, carno string, model models.Cars) error {
	collection := ctx.MongoDB.Collection("cars")

	filter := bson.D{
		primitive.E{
			Key:   "carno",
			Value: carno,
		},
	}

	update := bson.D{
		{"$set", bson.D{}},
	}

    if model.Brand != "" {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "brand", Value: model.Brand})
	}
	if model.Model != "" {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "model", Value: model.Model})
	}
	if model.Year != 0 {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "year", Value: model.Year})
	}
	if model.Price != 0 {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "price", Value: model.Price})
	}
	if model.IsNew || !model.IsNew {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "is_new", Value: model.IsNew})
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (c car) Delete(ctx *gofr.Context, carno string) (int, error) {
	collection := ctx.MongoDB.Collection("cars")

	filter := bson.D{
		primitive.E{
			Key:   "carno",
			Value: carno,
		},
	}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(res.DeletedCount), nil
}
