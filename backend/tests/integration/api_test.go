package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"user-management/cmd/server"
)

func TestAPI_CreateUser(t *testing.T) {
	app := server.SetupApp()

	t.Run("Valid user creation", func(t *testing.T) {
		payload := map[string]string{
			"name":     "John Doe",
			"email":    "john@example.com",
			"password": "password123",
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		app.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}

		var response map[string]string
		json.Unmarshal(rr.Body.Bytes(), &response)

		if _, exists := response["id"]; !exists {
			t.Errorf("expected id in response")
		}
	})

	t.Run("Invalid user creation", func(t *testing.T) {
		payload := map[string]string{
			"name":     "John Doe",
			"email":    "invalid-email",
			"password": "password123",
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		app.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}
