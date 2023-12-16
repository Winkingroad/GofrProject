package handlers_test

import (
	"github.com/Winkingroad/GofrProject/handlers"
	"github.com/Winkingroad/GofrProject/models"
	"github.com/Winkingroad/GofrProject/stores"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr/request"
	gofrLog "gofr.dev/pkg/log"
)

func newMock(t *testing.T) (gofrLog.Logger, *stores.MockCar) {  
	ctrl := gomock.NewController(t)  
	defer ctrl.Finish()  

	mockStore := stores.NewMockCar(ctrl)  
	mockLogger := gofrLog.NewMockLogger(io.Discard)  

	return mockLogger, mockStore  
}

func createContext(method string, params map[string]string, car interface{}, logger gofrLog.Logger, t *testing.T) *gofr.Context {  
	body, err := json.Marshal(car)  
	if err != nil {  
		t.Fatalf("Error while marshalling model: %v", err)  
	}  

	r := httptest.NewRequest(method, "/dummy/", bytes.NewBuffer(body))  
	query := r.URL.Query()

	for key, value := range params {  
		query.Add(key, value)  
	}  

	r.URL.RawQuery = query.Encode()  

	req := request.NewHTTPRequest(r)  

	return gofr.NewContext(nil, req, nil)  
}

func TestCreate(t *testing.T) {  
	mockLogger, mockStore := newMock(t)  
	h := handlers.New(mockStore)  
	car := models.Cars{  
		Brand: "Test Brand",
		Model: "Test Model",
		CarNo: "Test CarNo",
		Year:  2023,
		Price: 50000,
	}  

	expectedCar := []models.Cars{car}

	testCases := []struct {  
		desc      string  
		input     interface{}  
		mockCalls []*gomock.Call  
		expRes    interface{}  
		expErr    error  
	}{  
		{"success case", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return([]models.Cars{}, nil),
			mockStore.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), car).Return(nil),
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
		}, expectedCar, nil},  
		{"failure case - car already exists", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, &errors.Response{
				StatusCode: 200,
				Code:       "200",
				Reason:     "Car already exists",
			}),
		}, nil, &errors.Response{
			StatusCode: 200,
			Code:       "200",
			Reason:     "Car already exists",
		}},  
		{"failure case - error in creating car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return([]models.Cars{}, nil),
			mockStore.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), car).Return(errors.Error("database error")),
		}, nil, errors.Error("database error")},  
		{"failure case - error in getting car after creation", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return([]models.Cars{}, nil),
			mockStore.EXPECT().Create(gomock.AssignableToTypeOf(&gofr.Context{}), car).Return(nil),
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.Error("database error")),
		}, nil, errors.Error("database error")},
		{"failure case - Brand field is empty", models.Cars{Model: "Test Model", CarNo: "Test CarNo", Year: 2023, Price: 50000}, nil, nil, &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Brand field is required",
		}},
		{"failure case - Model field is empty", models.Cars{Brand: "Test Brand", CarNo: "Test CarNo", Year: 2023, Price: 50000}, nil, nil, &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Model field is required",
		}},  
		{"failure case - CarNo field is empty", models.Cars{Brand: "Test Brand", Model: "Test Model", Year: 2023, Price: 50000}, nil, nil, &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "CarNo field is required",
		}},
		{"failure case - CarNo field is empty", models.Cars{Brand: "Test Brand", Model: "Test Model", CarNo: "Test CarNo", Price: 50000}, nil, nil, &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Year field is required and must be greater than zero",
		}},
		{"failure case - CarNo field is empty", models.Cars{Brand: "Test Brand", Model: "Test Model", CarNo: "Test CarNo", Year: 2023}, nil, nil, &errors.Response{
			StatusCode: 400,
			Code:       "400",
			Reason:     "Price field is required and must be greater than zero",
		}},
		
	} 

	for i, tc := range testCases {  
		t.Run(tc.desc, func(t *testing.T) {  
			ctx := createContext(http.MethodPost, nil, tc.input, mockLogger, t)  
			res, err := h.Create(ctx)  

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)  
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)  
		})  
	}  
}
func TestGet(t *testing.T) {  
	mockLogger, mockStore := newMock(t)  
	h := handlers.New(mockStore)  
	car := models.Cars{  
		Brand: "Test Brand",
		Model: "Test Model",
		CarNo: "Test CarNo",
		Year:  2023,
		Price: 50000,
	}  

	expectedCar := []models.Cars{car}

	testCases := []struct {  
		desc      string  
		input     interface{}  
		mockCalls []*gomock.Call  
		expRes    interface{}  
		expErr    error  
	}{  
		{"success case", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
		}, expectedCar, nil},  
		{"failure case", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.EntityNotFound{Entity: "car", ID: car.CarNo}),
	    }, nil, errors.EntityNotFound{Entity: "car", ID: car.CarNo}}, 
		{"failure case - database error", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.Error("database error")),
		}, nil, errors.Error("database error")},  
		{"failure case - invalid car number", models.Cars{CarNo: "Invalid CarNo"}, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), "Invalid CarNo").Return(nil, errors.EntityNotFound{Entity: "car", ID: "Invalid CarNo"}),
		}, nil, errors.EntityNotFound{Entity: "car", ID: "Invalid CarNo"}},
	}   
		 

	for i, tc := range testCases {  
		t.Run(tc.desc, func(t *testing.T) {  
			carNo := car.CarNo
			if tc.desc == "failure case - invalid car number" {
				carNo = "Invalid CarNo"
			}
			ctx := createContext(http.MethodGet, map[string]string{"carno": carNo}, tc.input, mockLogger, t)  
			res, err := h.Get(ctx, carNo)  
	
			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)  
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)  
		})  
	}  
	 
}

func TestUpdate(t *testing.T) {  
	mockLogger, mockStore := newMock(t)  
	h := handlers.New(mockStore)  
	car := models.Cars{  
		Brand: "Test Brand",
		Model: "Test Model",
		CarNo: "Test CarNo",
		Year:  2023,
		Price: 50000,
	}  

	expectedCar := []models.Cars{car}

	testCases := []struct {  
		desc      string  
		input     interface{}  
		mockCalls []*gomock.Call  
		expRes    interface{}  
		expErr    error  
	}{  
		{"success case", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
			mockStore.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo, car).Return(nil),
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
		}, expectedCar, nil},  
		{"failure case - car not found", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, nil),
		}, nil, errors.EntityNotFound{Entity: "car", ID: car.CarNo}},  
		{"failure case - database error when getting car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.Error("database error")),
		}, nil, errors.Error("database error")},  
		{"failure case - database error when updating car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
			mockStore.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo, car).Return(errors.Error("database error")),
		}, nil, errors.Error("database error")},
		{"failure case - database error when getting updated car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
			mockStore.EXPECT().Update(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo, car).Return(nil),
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.Error("database error")),
		}, nil, errors.Error("database error")},

	}  

	for i, tc := range testCases {  
		t.Run(tc.desc, func(t *testing.T) {  
			ctx := createContext(http.MethodPut, nil, tc.input, mockLogger, t)  
			res, err := h.Update(ctx,car.CarNo)  

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)  
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)  
		})  
	}  
}

func TestDelete(t *testing.T) {  
	mockLogger, mockStore := newMock(t)  
	h := handlers.New(mockStore)  
	car := models.Cars{  
		Brand: "Test Brand",
		Model: "Test Model",
		CarNo: "BMW555",
		Year:  2023,
		Price: 50000,
	}  

	expectedCar := []models.Cars{car}

	testCases := []struct {  
		desc      string  
		input     interface{}  
		mockCalls []*gomock.Call  
		expRes    interface{}  
		expErr    error  
	}{  
		{"success case", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
			mockStore.EXPECT().Delete(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(1, nil),
		}, fmt.Sprintf("%v car Removed!", car.CarNo), nil},  
		{"failure case - car not found", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, nil),
		}, nil, errors.EntityNotFound{Entity: "car", ID: car.CarNo}},  
		{"failure case - database error when getting car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(nil, errors.Error("database error")),
		}, nil, errors.Error("database error")},  
		{"failure case - database error when deleting car", car, []*gomock.Call{  
			mockStore.EXPECT().Get(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(expectedCar, nil),
			mockStore.EXPECT().Delete(gomock.AssignableToTypeOf(&gofr.Context{}), car.CarNo).Return(0, errors.Error("database error")),
		}, nil, errors.Error("database error")},
	}


	for i, tc := range testCases {  
		t.Run(tc.desc, func(t *testing.T) {  
			ctx := createContext(http.MethodDelete, nil, tc.input, mockLogger, t)  
			res, err := h.Delete(ctx,car.CarNo)  

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)  
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)  
		})  
	}  
}
