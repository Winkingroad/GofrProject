package handlers_test

import (
	"testing"
	"io"
	"encoding/json"
	"net/http/httptest"
	"bytes"
	"ZopSmartproject/models"
	"ZopSmartproject/stores"
	"ZopSmartproject/handlers"
	"github.com/golang/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
	gofrLog "gofr.dev/pkg/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"fmt"
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
		// Add more test cases as needed
	}  

	for i, tc := range testCases {  
		t.Run(tc.desc, func(t *testing.T) {  
			ctx := createContext(http.MethodGet, map[string]string{"carno": car.CarNo}, tc.input, mockLogger, t)  
			res, err := h.Get(ctx, car.CarNo)  

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
