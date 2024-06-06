package task

import (
	"bytes"
	"cmd/main.go/types"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskHandlers(t *testing.T) {
	store := &mockTaskStore{tasks: make(map[string]types.Task)}
	handler := NewHandler(store)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	// POST
	t.Run("should handler create task", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Test Task",
			Description: "This is a test task",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
	t.Run("shouldn't handler create task if the title field is empty", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "",
			Description: "This is a test task",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	// GET
	t.Run("should handler get tasks", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should handler get task by ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks/89d9777c857a7fc95844", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("shouldn't handler get task with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks/invalid_ID", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler get task that don't exist", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks/77d77777c777a7fc7777", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})

	// PUT
	t.Run("should handler update task", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Updated Task Title",
			Description: "Updated task description",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("should handler update task if only description was given", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "",
			Description: "Updated only the description",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("shouldn't handler update task that don't exist", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Updated Task Title",
			Description: "Updated task description",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", "/tasks/77d77777c777a7fc7777", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})
	t.Run("shouldn't handler update task that don't exist", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Updated Task Title",
			Description: "Updated task description",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", "/tasks/invalid_ID", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	// PUT /{id}/complete
	t.Run("should handler complete task", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844/complete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("shouldn't handler complete task", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844/complete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler complete task with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/invalid_ID/complete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler complete task that don't exist", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/77d77777c777a7fc7777/complete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})

	// PUT /{id}/incomplete
	t.Run("should handler incomplete task", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844/incomplete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("shouldn't handler incomplete task", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/89d9777c857a7fc95844/incomplete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler incomplete task with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/invalid_ID/incomplete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler incomplete task that don't exist", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/tasks/77d77777c777a7fc7777/incomplete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})

	// DELETE
	t.Run("should handler delete task", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks/89d9777c857a7fc95844", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
	t.Run("shouldn't handler delete task with invalid ID", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks/invalid_ID", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("shouldn't handler delete task that don't exist", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks/77d77777c777a7fc7777", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, got %d", http.StatusNotFound, rr.Code)
		}
	})
}

type mockTaskStore struct {
	tasks map[string]types.Task
}

func (m *mockTaskStore) CreateTask(payload types.TaskPayLoad) error {
	id := "89d9777c857a7fc95844" // Mocked ID
	task := types.Task{
		ID:          id,
		Title:       payload.Title,
		Description: payload.Title,
		Status:      false,
	}
	m.tasks[id] = task
	return nil
}
func (m *mockTaskStore) ListTasks() ([]*types.Task, error) {
	tasks := make([]*types.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
func (m *mockTaskStore) GetTaskByID(id string) (*types.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}
func (m *mockTaskStore) UpdateTask(task types.Task) error {
	m.tasks[task.ID] = task
	return nil
}
func (m *mockTaskStore) DeleteTask(id string) (types.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return task, errors.New("task not found")
	}
	delete(m.tasks, id)
	return task, nil
}
