package stores

import (
    "github.com/Winkingroad/GofrProject/models"
    "gofr.dev/pkg/gofr"
)

type Car interface {
    Get(ctx *gofr.Context, carno string) ([]models.Cars, error)
    Create(ctx *gofr.Context, model models.Cars) error
    Update(ctx *gofr.Context, carno string, model models.Cars) error
    Delete(ctx *gofr.Context, carno string) (int, error)
    GetAllCars(ctx *gofr.Context) ([]models.Cars, error) 
}
