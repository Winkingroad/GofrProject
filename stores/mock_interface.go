// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_stores is a generated GoMock package.
package stores

import (
	models "ZopSmartproject/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockCar is a mock of Car interface.
type MockCar struct {
	ctrl     *gomock.Controller
	recorder *MockCarMockRecorder
}

// MockCarMockRecorder is the mock recorder for MockCar.
type MockCarMockRecorder struct {
	mock *MockCar
}

// NewMockCar creates a new mock instance.
func NewMockCar(ctrl *gomock.Controller) *MockCar {
	mock := &MockCar{ctrl: ctrl}
	mock.recorder = &MockCarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCar) EXPECT() *MockCarMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCar) Create(ctx *gofr.Context, model models.Cars) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCarMockRecorder) Create(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCar)(nil).Create), ctx, model)
}

// Delete mocks base method.
func (m *MockCar) Delete(ctx *gofr.Context, carno string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, carno)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCarMockRecorder) Delete(ctx, carno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCar)(nil).Delete), ctx, carno)
}

// Get mocks base method.
func (m *MockCar) Get(ctx *gofr.Context, carno string) ([]models.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, carno)
	ret0, _ := ret[0].([]models.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCarMockRecorder) Get(ctx, carno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCar)(nil).Get), ctx, carno)
}

// GetAllCars mocks base method.
func (m *MockCar) GetAllCars(ctx *gofr.Context) ([]models.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCars", ctx)
	ret0, _ := ret[0].([]models.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCars indicates an expected call of GetAllCars.
func (mr *MockCarMockRecorder) GetAllCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCars", reflect.TypeOf((*MockCar)(nil).GetAllCars), ctx)
}

// Update mocks base method.
func (m *MockCar) Update(ctx *gofr.Context, carno string, model models.Cars) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, carno, model)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCarMockRecorder) Update(ctx, carno, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCar)(nil).Update), ctx, carno, model)
}