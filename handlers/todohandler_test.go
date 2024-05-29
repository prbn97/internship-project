package handlers

import (
	"api/main.go/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupEnv(t *testing.T) (*TodoHandler, func(), []byte, models.Todo, []byte) {
	todoHandler := NewTodoHandler()

	initialTodos := []models.Todo{
		{Title: "toDo #1", Description: "Description field", Completed: true},
		{Title: "toDo #2", Description: "Description field"},
	}

	var todoItem models.Todo

	for _, todo := range initialTodos {
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatalf("error marshalling todo: %v", err)
		}

		postResponse := httptest.NewRecorder()
		postRequest := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(todoJSON))
		todoHandler.Create(postResponse, postRequest)

		response := postResponse.Result()
		defer response.Body.Close()
		if response.StatusCode != http.StatusCreated {
			t.Fatalf("failed to create initial todo, status code: %d", response.StatusCode)
		}

		var createdTodo models.Todo
		if err := json.NewDecoder(response.Body).Decode(&createdTodo); err != nil {
			t.Fatalf("error decoding response body: %v", err)
		}
		todoItem = createdTodo
	}

	teardown := func() {
		todoHandler.store.Lock()
		defer todoHandler.store.Unlock()
		for key := range todoHandler.store.m {
			delete(todoHandler.store.m, key)
		}
	}

	NewTodo := models.Todo{
		Title:       "toDo #3",
		Description: "Description field",
		Completed:   false,
	}
	NewTodoJSON, _ := json.Marshal(NewTodo)

	todoData := models.Todo{
		Description: "Updated Description",
		Completed:   true,
	}
	UpTodoJSON, _ := json.Marshal(todoData)

	return todoHandler, teardown, NewTodoJSON, todoItem, UpTodoJSON
}

func TestTodoHandler_Create(t *testing.T) {
	todoHandler, teardown, todoJSON, _, _ := setupEnv(t)
	defer teardown()

	tests := []struct {
		name         string
		input        []byte
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "ValidTodo",
			input:        todoJSON,
			expectedCode: http.StatusCreated,
		},
		{
			name:         "InvalidJSON",
			input:        []byte(`{"title": "todo", "description": `), // malformed JSON
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Request body contains badly-formed JSON",
		},
		{
			name:         "UnknownField",
			input:        []byte(`{"title": "todo", "unknown": "field"}`),
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "Request body contains unknown field \"unknown\"",
		},
		{
			name:         "EmptyTitle",
			input:        []byte(`{"description": "no title"}`),
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "title is required",
		},
		{
			name:         "ProvidedID",
			input:        []byte(`{"id": "123", "title": "todo", "description": "no title"}`),
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "ID cannot be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRequest := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(tt.input))
			postResponse := httptest.NewRecorder()

			todoHandler.Create(postResponse, postRequest)

			response := postResponse.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}

			if tt.expectedMsg != "" {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					t.Errorf("expected error to be nil, got %v", err)
				}
				var errorMsg map[string]string
				if err := json.Unmarshal(body, &errorMsg); err != nil {
					t.Fatalf("error unmarshalling response body: %v", err)
				}
				if errorMsg["message"] != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, errorMsg["message"])
				}
			}
		})
	}
}

func TestTodoHandler_List(t *testing.T) {
	todoHandler, teardown, _, _, _ := setupEnv(t)
	defer teardown()

	getRequest := httptest.NewRequest(http.MethodGet, "/todos", nil)
	getResponse := httptest.NewRecorder()

	todoHandler.List(getResponse, getRequest)

	response := getResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	var todos []models.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		t.Fatalf("error unmarshalling response body: %v", err)
	}

	if len(todos) != 2 {
		t.Errorf("expected 2 todos items, got %d", len(todos))
	}
}

func TestTodoHandler_Get(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "ValidGet",
			url:          "/todos/" + createdTodo.ID,
			expectedCode: http.StatusOK,
		},
		{
			name:         "InvalidIDFormat",
			url:          "/todos/invalidID",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid ID",
		},
		{
			name:         "TodoNotFound",
			url:          "/todos/00000000000000000000", // Assuming this ID doesn't exist
			expectedCode: http.StatusNotFound,
			expectedMsg:  "ID not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getRequest := httptest.NewRequest(http.MethodGet, tt.url, nil)
			getResponse := httptest.NewRecorder()

			todoHandler.Get(getResponse, getRequest)

			response := getResponse.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}

			if tt.expectedMsg != "" {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					t.Errorf("expected error to be nil, got %v", err)
				}
				var errorMsg map[string]string
				if err := json.Unmarshal(body, &errorMsg); err != nil {
					t.Fatalf("error unmarshalling response body: %v", err)
				}
				if errorMsg["message"] != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, errorMsg["message"])
				}
			}
		})
	}
}

func TestTodoHandler_Update(t *testing.T) {
	todoHandler, teardown, _, createdTodo, UpTodoJSON := setupEnv(t)
	defer teardown()

	tests := []struct {
		name         string
		input        []byte
		url          string
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "ValidUpdate",
			input:        UpTodoJSON,
			url:          "/todos/" + createdTodo.ID,
			expectedCode: http.StatusOK,
		},
		{
			name:         "InvalidID",
			input:        UpTodoJSON,
			url:          "/todos/invalidID",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid ID",
		},
		{
			name:         "TodoNotFound",
			input:        UpTodoJSON,
			url:          "/todos/00000000000000000000", // Assuming this ID doesn't exist
			expectedCode: http.StatusNotFound,
			expectedMsg:  "ID not found",
		},
		{
			name:         "IDInPayload",
			input:        []byte(`{"id": "newID", "title": "updated title"}`),
			url:          "/todos/" + createdTodo.ID,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "ID cannot be updated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			putRequest := httptest.NewRequest(http.MethodPut, tt.url, bytes.NewReader(tt.input))
			putResponse := httptest.NewRecorder()

			todoHandler.Update(putResponse, putRequest)

			response := putResponse.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}

			if tt.expectedMsg != "" {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					t.Errorf("expected error to be nil, got %v", err)
				}
				var errorMsg map[string]string
				if err := json.Unmarshal(body, &errorMsg); err != nil {
					t.Fatalf("error unmarshalling response body: %v", err)
				}
				if errorMsg["message"] != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, errorMsg["message"])
				}
			}
		})
	}
}

func TestTodoHandler_Delete(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "ValidDelete",
			url:          "/todos/" + createdTodo.ID,
			expectedCode: http.StatusOK,
		},
		{
			name:         "InvalidIDFormat",
			url:          "/todos/invalidID",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid ID",
		},
		{
			name:         "TodoNotFound",
			url:          "/todos/00000000000000000000", // Assuming this ID doesn't exist
			expectedCode: http.StatusNotFound,
			expectedMsg:  "ID not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delRequest := httptest.NewRequest(http.MethodDelete, tt.url, nil)
			delResponse := httptest.NewRecorder()

			todoHandler.Delete(delResponse, delRequest)

			response := delResponse.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}

			if tt.expectedMsg != "" {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					t.Errorf("expected error to be nil, got %v", err)
				}
				var errorMsg map[string]string
				if err := json.Unmarshal(body, &errorMsg); err != nil {
					t.Fatalf("error unmarshalling response body: %v", err)
				}
				if errorMsg["message"] != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, errorMsg["message"])
				}
			}
		})
	}
}

func TestTodoHandler_MarkComplete(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedMsg  string
	}{
		{
			name:         "ValidMarkComplete",
			url:          "/todos/" + createdTodo.ID + "/complete",
			expectedCode: http.StatusOK,
		},
		{
			name:         "InvalidIDFormat",
			url:          "/todos/invalidID/complete",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid ID",
		},
		{
			name:         "TodoNotFound",
			url:          "/todos/00000000000000000000/complete", // Assuming this ID doesn't exist
			expectedCode: http.StatusNotFound,
			expectedMsg:  "ID not found",
		},
		{
			name:         "AlreadyCompleted",
			url:          "/todos/" + createdTodo.ID + "/complete",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "the task is already done",
		},
	}

	// Mark the initial todo as completed for the "AlreadyCompleted" test case
	todoHandler.store.Lock()
	createdTodo.Completed = false
	todoHandler.store.m[createdTodo.ID] = createdTodo
	todoHandler.store.Unlock()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			completeRequest := httptest.NewRequest(http.MethodPut, tt.url, nil)
			completeResponse := httptest.NewRecorder()

			todoHandler.MarkComplete(completeResponse, completeRequest)

			response := completeResponse.Result()
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}

			if tt.expectedMsg != "" {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					t.Errorf("expected error to be nil, got %v", err)
				}
				var errorMsg map[string]string
				if err := json.Unmarshal(body, &errorMsg); err != nil {
					t.Fatalf("error unmarshalling response body: %v", err)
				}
				if errorMsg["message"] != tt.expectedMsg {
					t.Errorf("expected error message %q, got %q", tt.expectedMsg, errorMsg["message"])
				}
			}
		})
	}
}
