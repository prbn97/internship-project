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

	postRequest := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(todoJSON))
	postResponse := httptest.NewRecorder()

	todoHandler.Create(postResponse, postRequest)

	response := postResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, response.StatusCode)
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

	getRequest := httptest.NewRequest(http.MethodGet, "/todos/"+createdTodo.ID, nil)
	getResponse := httptest.NewRecorder()

	todoHandler.Get(getResponse, getRequest)

	response := getResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

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

	putRequest := httptest.NewRequest(http.MethodPut, "/todos/"+createdTodo.ID, bytes.NewReader(UpTodoJSON))
	putResponse := httptest.NewRecorder()

	todoHandler.Update(putResponse, putRequest)

	response := putResponse.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}
}

func TestTodoHandler_Delete(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	delRequest := httptest.NewRequest(http.MethodDelete, "/todos/"+createdTodo.ID, nil)
	delResponse := httptest.NewRecorder()

	todoHandler.Delete(delResponse, delRequest)

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

func TestTodoHandler_MarkComplete(t *testing.T) {
	todoHandler, teardown, _, createdTodo, _ := setupEnv(t)
	defer teardown()

	completeRequest := httptest.NewRequest(http.MethodPut, "/todos/"+createdTodo.ID+"/complete", nil)
	completeResponse := httptest.NewRecorder()

	todoHandler.MarkComplete(completeResponse, completeRequest)

	response := completeResponse.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(response.Body).Decode(&updatedTodo); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}
	if !updatedTodo.Completed {
		t.Errorf("expected todo to be marked as completed")
	}
}
