package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	data := map[string]string{"message": "test"}

	err := WriteJSON(rec, http.StatusOK, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	if contentType := rec.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", contentType)
	}

	var response map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error decoding response body: %v", err)
	}

	if response["message"] != "test" {
		t.Errorf("expected message 'test', got %s", response["message"])
	}
}

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()
	errMessage := "something went wrong"

	WriteError(rec, http.StatusBadRequest, errors.New(errMessage))

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, status)
	}

	var response map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error decoding response body: %v", err)
	}

	if response["error"] != errMessage {
		t.Errorf("expected error message '%s', got %s", errMessage, response["error"])
	}
}

func TestParseJSON(t *testing.T) {
	type Payload struct {
		Name string `json:"name"`
	}

	t.Run("valid JSON", func(t *testing.T) {
		body := `{"name":"test"}`
		req := httptest.NewRequest("POST", "/", strings.NewReader(body))
		var payload Payload

		err := ParseJSON(req, &payload)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if payload.Name != "test" {
			t.Errorf("expected name 'test', got %s", payload.Name)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		body := `{"name":"test"`
		req := httptest.NewRequest("POST", "/", strings.NewReader(body))
		var payload Payload

		err := ParseJSON(req, &payload)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		var payload Payload

		err := ParseJSON(req, &payload)
		if err == nil || err.Error() != "empty request body" {
			t.Fatalf("expected error 'empty request body', got: %v", err)
		}
	})
}

func TestGetTokenFromRequest(t *testing.T) {
	t.Run("token from Authorization header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer test_token")

		token := GetTokenFromRequest(req)
		if token != "Bearer test_token" {
			t.Errorf("expected token 'Bearer test_token', got %s", token)
		}
	})

	t.Run("token from query parameter", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?token=test_token", nil)

		token := GetTokenFromRequest(req)
		if token != "test_token" {
			t.Errorf("expected token 'test_token', got %s", token)
		}
	})

	t.Run("no token provided", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)

		token := GetTokenFromRequest(req)
		if token != "" {
			t.Errorf("expected empty token, got %s", token)
		}
	})
}
