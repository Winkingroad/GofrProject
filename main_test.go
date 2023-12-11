package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(5 * time.Second)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"register success", http.MethodPost, "signup", http.StatusCreated, []byte(`{"username": "test2","password": "1234567890"}`)},
		{"login success", http.MethodPost, "login", http.StatusCreated, []byte(`{"username": "test2","password": "1234567890"}`)},
		{"get success", http.MethodGet, "cars", http.StatusOK, nil},
		{"create success", http.MethodPost, "cars", http.StatusOK, []byte(`{"brand": "BMW","carno": "BMW-555-KH","model": "X5","year": 2021,"price": 45000,"is_new": true}`)},
		{"get success", http.MethodGet, "cars/BMW-555-KH", http.StatusOK, nil},
		{"update success", http.MethodPut, "cars/BMW-555-KH", http.StatusOK, []byte(`{"brand": "BMW","model": "X0","year": 2024}`)},
		{"delete success", http.MethodDelete, "cars/BMW-555-KH", http.StatusOK, nil},
		{"get unknown endpoint", http.MethodGet, "car", http.StatusNotFound, nil},
		{"get unknown car", http.MethodGet, "cars/BMW-555", http.StatusNotFound, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:8097/"+tc.endpoint, bytes.NewBuffer(tc.body))
		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue // move to next test
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}
