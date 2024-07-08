package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	types "github.com/prbn97/internship-project/services/models"
)

type mockUserStore struct{}

func TestUsersHandler(t *testing.T) { // func to test all end points
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	t.Run("Login should fail, login user - invalid payload", func(t *testing.T) {
		payload := types.LoginUserPayload{
			Email:    "pauloreis.email.com",
			Password: "segredo",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Register should fail register user - invalid payload", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Paulo",
			LastName:  "Reis",
			Email:     "pauloreis.email.com",
			Password:  "segredo",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	// more tests needed!
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) UpdateUser(u types.User) error {
	return nil
}
