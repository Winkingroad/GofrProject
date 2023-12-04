package stores

import (
	"ZopSmartproject/models"
	"gofr.dev/pkg/gofr"
)

type Car interface {
	Get(ctx *gofr.Context, id string) ([]models.Cars, error)
	Create(ctx *gofr.Context, model models.Cars) error
	Update(ctx *gofr.Context, id string, model models.Cars) error
	Delete(ctx *gofr.Context, id string) (int, error)
}