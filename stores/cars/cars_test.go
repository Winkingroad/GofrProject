package cars

import (
	"context"
	"testing"

	"github.com/Winkingroad/GofrProject/models"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/gofr"

	"github.com/stretchr/testify/assert"
)

func initializeTest(t *testing.T) *gofr.Context {
	app := gofr.New()

	// initializing the seeder
	seeder := datastore.NewSeeder(&app.DataStore, "../../db")
	seeder.RefreshMongoCollections(t, "cars")

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	return ctx
}

func TestCustomer_Get(t *testing.T) {
	tests := []struct {
		desc string
		CarNo string
		resp []models.Cars
		err  error
	}{
		{"get single entity", "HON456", []models.Cars{{Brand: "Honda",
        CarNo: "HON456",
        Model: "Civic",
        Year: 2019,
        Price: 22000,
        IsNew: false}}, nil},
		{"get unregistered entity", "HON123", []models.Cars{}, nil},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		resp, err := store.Get(ctx, tc.CarNo)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestCustomer_GetAllCars(t *testing.T) {
	expectedCars := []models.Cars{
		{
			Brand:  "Ford",
			CarNo:  "FRD789",
			Model:  "Mustang",
			Year:   2022,
			Price:  35000,
			IsNew:  true,
		},
		{
			Brand:  "Honda",
			CarNo:  "HON456",
			Model:  "Civic",
			Year:   2019,
			Price:  22000,
			IsNew:  false,
		},
		{
			Brand:  "BMW",
			CarNo:  "BMW555",
			Model:  "X5",
			Year:   2021,
			Price:  45000,
			IsNew:  true,
		},
	}

	tests := []struct {
		desc string
		resp []models.Cars
		err  error
	}{
		{
			desc: "get all cars",
			resp: expectedCars,
			err:  nil,
		},
	}

	store := New()
	ctx := initializeTest(t)

	for i, tc := range tests {
		resp, err := store.GetAllCars(ctx)

		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestCustomer_Create(t *testing.T) {
	newCar := models.Cars{
		Brand:  "Audi",
		CarNo:  "AUD123",
		Model:  "A4",
		Year:   2023,
		Price:  50000,
		IsNew:  true,
	}

	store := New()
	ctx := initializeTest(t)

	err := store.Create(ctx, newCar)

	assert.NoError(t, err, "Failed to create a new car")


}

func TestCustomer_Update(t *testing.T) {
	updatedCar := models.Cars{
		Brand:  "BMW",
		CarNo:  "BMW555",
		Model:  "X6",
		Year:   2023,
		Price:  55000,
		IsNew:  true,
	}

	store := New()
	ctx := initializeTest(t)

	err := store.Update(ctx, "BMW555", updatedCar)

	assert.NoError(t, err, "Failed to update the car")

	
}

func TestCustomer_Delete(t *testing.T) {
	store := New()
	ctx := initializeTest(t)

	countBeforeDeletion := 4

	deletedCount, err := store.Delete(ctx, "BMW555")

	assert.NoError(t, err, "Failed to delete the car")
	assert.Equal(t, 1, deletedCount, "Expected one car to be deleted")

	countAfterDeletion := 3

	assert.Equal(t, countBeforeDeletion-1, countAfterDeletion, "Incorrect count after deletion")

	
}

