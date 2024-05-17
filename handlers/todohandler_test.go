package handlers

import (
	"api/main.go/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupEnv(t *testing.T) (*TodoHandler, func(), []byte, models.Todo, []byte) {
	// create todoHandler and toDos item
	todoHandler := NewTodoHandler()

	initialTodos := []models.Todo{
		{Title: "toDo #1", Description: "Description field", Completed: true},
		{Title: "toDo #2", Description: "Description field"},
	}

	var todoItem models.Todo

	// add initial Todos items in handler
	for _, todo := range initialTodos {
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			t.Fatalf("error marshalling todo: %v", err)
		}

		postResponse := httptest.NewRecorder()
		postRequest := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(todoJSON))
		todoHandler.ServeHTTP(postResponse, postRequest)

		response := postResponse.Result()
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			t.Fatalf("failed to create initial todo, status code: %d", response.StatusCode)
		}

		var createdTodo models.Todo
		if err := json.NewDecoder(response.Body).Decode(&createdTodo); err != nil {
			t.Fatalf("error decoding response body: %v", err)
		}
		// Store the ID of the last created Todo
		todoItem = createdTodo
	}

	// teardown for clear the todoHandler
	teardown := func() {
		todoHandler.store.Lock()
		defer todoHandler.store.Unlock()
		for key := range todoHandler.store.m {
			delete(todoHandler.store.m, key)
		}
	}

	// New toDo
	NewTodo := models.Todo{
		Title:       "toDo #3",
		Description: "Description field",
		Completed:   false,
	}
	NewTodoJSON, _ := json.Marshal(NewTodo)

	// Prepare toDo updated data
	todoData := models.Todo{
		Description: "Updated Description",
		Completed:   true,
	}
	UpTodoJSON, _ := json.Marshal(todoData)

	return todoHandler, teardown, NewTodoJSON, todoItem, UpTodoJSON
}

func TestTodoHandler_List(t *testing.T) {
	todoHandler, teardown, _, _, _ := setupEnv(t)
	defer teardown()

	// Get request
	getRequest := httptest.NewRequest(http.MethodGet, "/todos", nil)
	getResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler.ServeHTTP)

	handler.ServeHTTP(getResponse, getRequest)

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

func TestTodoHandler_Create(t *testing.T) {
	todoHandler, teardown, todoJSON, _, _ := setupEnv(t)
	defer teardown()

	// POST request (toDo #3)
	postRequest := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(todoJSON))
	postResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler.ServeHTTP)

	handler.ServeHTTP(postResponse, postRequest)

	response := postResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}

func TestTodoHandler_Get(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	// GET request by id
	getRequest := httptest.NewRequest(http.MethodGet, "/todos/"+createdTodo.ID, nil)
	getResponse := httptest.NewRecorder()
	todoHandler.ServeHTTP(getResponse, getRequest)
	response := getResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Validate get response
	var retrievedTodo models.Todo
	if err := json.NewDecoder(response.Body).Decode(&retrievedTodo); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}
	if !reflect.DeepEqual(createdTodo, retrievedTodo) {
		t.Fatalf("expected todo %+v, got %+v", createdTodo, retrievedTodo)
	}
}

func TestTodoHandler_Update(t *testing.T) {
	todoHandler, teardown, _, createdTodo, UpTodoJSON := setupEnv(t)
	defer teardown()

	// UPDATE by ID
	putRequest := httptest.NewRequest(http.MethodPut, "/todos/"+createdTodo.ID, bytes.NewReader(UpTodoJSON))
	putResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(todoHandler.ServeHTTP)

	handler.ServeHTTP(putResponse, putRequest)

	response := putResponse.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}

func TestTodoHandler_Delete(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	// DELETE by ID
	delRequest := httptest.NewRequest(http.MethodDelete, "/todos/"+createdTodo.ID, nil)
	delResponse := httptest.NewRecorder()
	todoHandler.ServeHTTP(delResponse, delRequest)

	response := delResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	_, exists := todoHandler.store.m[createdTodo.ID]
	if exists {
		t.Fatalf("todo was not deleted successfully")
	}
}
