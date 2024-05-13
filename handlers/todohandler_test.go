package handlers

import (
	"api/main.go/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTodoHandler_List(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(NewTodoHandler().ServeHTTP)

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
}

func TestTodoHandler_Create(t *testing.T) {

	todo := []byte(`{"title":"Test Todo","description":"Test Description","completed":false}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(todo))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(NewTodoHandler().ServeHTTP)

	handler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

}

func TestTodoHandler_Get(t *testing.T) {
	// Criar um manipulador de Todo
	todoHandler := NewTodoHandler()

	newTodo := models.Todo{
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
	}
	todoJSON, err := json.Marshal(newTodo)
	if err != nil {
		t.Fatalf("error marshalling todo: %v", err)
	}

	reqPost := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(todoJSON))

	wPost := httptest.NewRecorder()
	todoHandler.ServeHTTP(wPost, reqPost)
	resPost := wPost.Result()
	defer resPost.Body.Close()
	if resPost.StatusCode != http.StatusOK {
		t.Fatalf("failed to create todo, status code: %d", resPost.StatusCode)
	}

	var createdTodo models.Todo
	if err := json.NewDecoder(resPost.Body).Decode(&createdTodo); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	// Get request by id
	reqGet := httptest.NewRequest(http.MethodGet, "/todos/"+createdTodo.ID, nil)
	wGet := httptest.NewRecorder()
	todoHandler.ServeHTTP(wGet, reqGet)
	resGet := wGet.Result()
	defer resGet.Body.Close()

	if resGet.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resGet.StatusCode)
	}

	// Validate get response
	var retrievedTodo models.Todo
	if err := json.NewDecoder(resGet.Body).Decode(&retrievedTodo); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}
	if !reflect.DeepEqual(createdTodo, retrievedTodo) {
		t.Fatalf("expected todo %+v, got %+v", createdTodo, retrievedTodo)
	}
}

func TestTodoHandler_Delete(t *testing.T) {

	todoHandler := NewTodoHandler()

	newTodo := models.Todo{
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
	}
	todoJSON, err := json.Marshal(newTodo)
	if err != nil {
		t.Fatalf("error marshalling todo: %v", err)
	}

	reqPost := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(todoJSON))
	wPost := httptest.NewRecorder()
	todoHandler.ServeHTTP(wPost, reqPost)
	resPost := wPost.Result()
	defer resPost.Body.Close()
	if resPost.StatusCode != http.StatusOK {
		t.Fatalf("failed to create todo, status code: %d", resPost.StatusCode)
	}

	var createdTodo models.Todo
	if err := json.NewDecoder(resPost.Body).Decode(&createdTodo); err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	// Delete request by id
	reqDelete := httptest.NewRequest(http.MethodDelete, "/todos/"+createdTodo.ID, nil)
	wDelete := httptest.NewRecorder()
	todoHandler.ServeHTTP(wDelete, reqDelete)
	resDelete := wDelete.Result()
	defer resDelete.Body.Close()

	if resDelete.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resDelete.StatusCode)
	}

	_, exists := todoHandler.store.m[createdTodo.ID]
	if exists {
		t.Fatalf("todo was not deleted successfully")
	}
}
