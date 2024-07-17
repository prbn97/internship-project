package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	types "github.com/prbn97/internship-project/services/models"
	"gopkg.in/go-playground/assert.v1"
)

// mockUserStore is an in-memory store for users, used for testing
type mockUserStore struct {
	users map[string]types.User
}

// MockUserStore initializes the mockUserStore with a default user
func MockUserStore() *mockUserStore {
	store := &mockUserStore{
		users: make(map[string]types.User),
	}

	// Adding a default user
	store.users["usuario01@gmail.com"] = types.User{
		ID:        1,
		FirstName: "Default",
		LastName:  "User",
		Email:     "usuario01@gmail.com",
		Password:  "$2a$10$slORHDfJMaKuSLONrn9h5eCUjw5NL5BW2x4xeYZwilnJsz1.Vz02C", // "segredo" hashed
	}

	return store
}

func TestUsersHandler(t *testing.T) { // func to test all end points
	userStore := MockUserStore()
	handler := NewHandler(userStore)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	t.Run("Register should succeed, correct payload", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Paulo",
			LastName:  "Reis",
			Email:     "pauloreis@gmail.com", // valid email
			Password:  "senhafraca",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})
	t.Run("Register should fail, user invalid payload", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "Paulo",
			LastName:  "Reis",
			Email:     "pauloreis.email.com", // invalid email
			Password:  "senhafraca",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
	t.Run("Register should fail, user already exist!", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "already",
			LastName:  "exist",
			Email:     "usuario01@gmail.com", // default email
			Password:  "senhafraca",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("Login succeed, correct payload", func(t *testing.T) {
		payload := types.LoginUserPayload{
			Email:    "usuario01@gmail.com", // default email
			Password: "segredo",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("Login should fail, user invalid payload", func(t *testing.T) {
		payload := types.LoginUserPayload{
			Email:    "pauloreis.email.com", // invalid email
			Password: "segredo",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
	t.Run("Login should fail, incorrect payload", func(t *testing.T) {
		payload := types.LoginUserPayload{
			Email:    "usuario01@gmail.com", // default email
			Password: "senha-errada",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

// Implement mock methods for the UserStore interface
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	user, exists := m.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (m *mockUserStore) CreateUser(u types.User) error {
	if _, exists := m.users[u.Email]; exists {
		return errors.New("user already exists")
	}
	m.users[u.Email] = u
	return nil
}

func (m *mockUserStore) UpdateUser(u types.User) error {
	if _, exists := m.users[u.Email]; !exists {
		return errors.New("user not found")
	}
	m.users[u.Email] = u
	return nil
}
