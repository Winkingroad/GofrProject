package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)
	var authToken string

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
		{"create success", http.MethodPost, "cars", http.StatusCreated, []byte(`{"brand": "BMW","carno": "BMW-555-KH","model": "X5","year": 2021,"price": 45000,"is_new": true}`)},
		{"get success", http.MethodGet, "cars/BMW-555-KH", http.StatusOK, nil},
		{"update success", http.MethodPut, "cars/BMW-555-KH", http.StatusOK, []byte(`{"brand": "BMW","model": "X0","year": 2024}`)},
		{"delete success", http.MethodDelete, "cars/BMW-555-KH", http.StatusNoContent, nil},
		{"get unknown endpoint", http.MethodGet, "car", http.StatusNotFound, nil},
		{"get unknown car", http.MethodGet, "cars/BMW-555", http.StatusNotFound, nil},
	}

	for i, tc := range tests {
        req, _ := request.NewMock(tc.method, "http://localhost:9000/"+tc.endpoint, bytes.NewBuffer(tc.body))
        if authToken != "" {
            req.Header.Set("Authorization", "Bearer "+authToken)
        }

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            t.Fatalf("Failed to make request for test %d: %v", i, err)
        }
        defer resp.Body.Close()

       
        if resp.StatusCode != tc.statusCode {
            t.Errorf("Test %d: Expected status code %d, got %d", i, tc.statusCode, resp.StatusCode)
        }

      
        if tc.desc == "login success" && resp.StatusCode == http.StatusCreated {
            var response map[string]interface{}
            err := json.NewDecoder(resp.Body).Decode(&response)
            if err != nil {
                t.Errorf("Failed to parse login response body: %v", err)
                continue
            }
            data, exists := response["data"].(map[string]interface{})
            if !exists {
                t.Errorf("Data not found in login response for test %d", i)
                continue
            }
            token, exists := data["token"].(string)
            if !exists {
                t.Errorf("Token not found in login response for test %d", i)
                continue
            }
            authToken = token
        }
    }
}