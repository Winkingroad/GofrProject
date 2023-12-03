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

func (c car) Get(ctx *gofr.Context, id string) ([]models.Cars, error) {
	 resp := make([]models.Cars, 0)

	 collection := ctx.MongoDB.Collection("cars")

	 filter := bson.D{}

	 if id != "" {
		idFiltter := primitive.E{
			Key: "id",
			Value: id,
		}
		filter = append(filter, idFiltter)

	 }

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

	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (c car) Delete(ctx *gofr.Context, id string) (int, error) {
	collection := ctx.MongoDB.Collection("cars")

	filter := bson.D{
		primitive.E{
			Key: "id",
			Value: id,
		},
	}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, errors.DB{Err: err}
	}

	return int(res.DeletedCount), nil
}