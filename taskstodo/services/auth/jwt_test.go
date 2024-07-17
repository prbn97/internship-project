package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prbn97/internship-project/configs"
	types "github.com/prbn97/internship-project/services/models"
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

func TestCreateJWT(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}

func TestWithJWTAuth(t *testing.T) {
	mockUserStore := MockUserStore()

	// Create a handler for testing
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		// Check if the user ID is correctly set in the context
		userID := GetUserIDFromContext(r.Context())
		if userID != 1 {
			t.Errorf("expected userID 1, got %d", userID)
		}
	}

	// Generate a valid JWT token for testing
	token, err := CreateJWT([]byte(configs.Envs.JWTSecret), 1)
	if err != nil {
		t.Fatalf("error creating JWT: %v", err)
	}

	// Wrap the test handler with the JWT authentication middleware
	handler := WithJWTAuth(testHandler, mockUserStore)

	// Simulate a request with a valid JWT token
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestValidateJWT(t *testing.T) {
	// Generate a valid JWT token for testing
	token, err := CreateJWT([]byte(configs.Envs.JWTSecret), 1)
	if err != nil {
		t.Fatalf("error creating JWT: %v", err)
	}

	// Call validateJWT with the token
	_, err = validateJWT(token)
	if err != nil {
		t.Errorf("error validating JWT: %v", err)
	}
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	// Simulate an invalid JWT token
	invalidToken := "invalid_token_string"

	// Call validateJWT with the invalid token
	_, err := validateJWT(invalidToken)
	if err == nil {
		t.Error("expected error validating invalid JWT token, got nil")
	}
}

func TestPermissionDenied(t *testing.T) {
	// Mock a response writer
	w := httptest.NewRecorder()

	// Call permissionDenied
	permissionDenied(w)

	// Check the response status code
	if w.Code != http.StatusForbidden {
		t.Errorf("expected status code %d, got %d", http.StatusForbidden, w.Code)
	}
}

func TestGetUserIDFromContext(t *testing.T) {
	// Create a mock context with userID set
	ctx := context.WithValue(context.Background(), UserKey, 1)

	// Call GetUserIDFromContext
	userID := GetUserIDFromContext(ctx)
	if userID != 1 {
		t.Errorf("expected userID 1, got %d", userID)
	}
}

func TestGetUserIDFromContext_NoUserID(t *testing.T) {
	// Create an empty context
	ctx := context.Background()

	// Call GetUserIDFromContext
	userID := GetUserIDFromContext(ctx)
	if userID != -1 {
		t.Errorf("expected userID -1 (not found), got %d", userID)
	}
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
