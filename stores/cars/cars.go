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

func (c car) Update(ctx *gofr.Context, id string, model models.Cars) error {
	collection := ctx.MongoDB.Collection("cars")

	filter := bson.D{
		primitive.E{
			Key:   "id",
			Value: id,
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
	if model.IsNew {
		update[0].Value = append(update[0].Value.(bson.D), primitive.E{Key: "is_new", Value: model.IsNew})
	}
	
	_, err := collection.UpdateOne(ctx, filter, update)
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