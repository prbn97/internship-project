package tasks

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prbn97/internship-project/configs"
	"github.com/prbn97/internship-project/services/auth"
	types "github.com/prbn97/internship-project/services/models"
	"github.com/stretchr/testify/assert"
)

// mockUserStore is an in-memory store for users, used for testing
type mockUserStore struct {
	users map[string]types.User
}

// MockTaskStore é uma store em memória para tarefas, usada para testes
type mockTaskStore struct {
	tasks map[int]types.Task
}

// Mock User Storage with a default user
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

// MockTaskStore inicializa a MockTaskStore com uma tarefa padrão
func MockTaskStore() *mockTaskStore {
	store := &mockTaskStore{
		tasks: make(map[int]types.Task),
	}

	// Adicionando uma tarefa padrão
	store.tasks[1] = types.Task{
		ID:          1,
		Title:       "Default Task",
		Description: "This is a default task",
		UserID:      1,
		Status:      "ToDo",
	}

	return store
}

func TestHandler(t *testing.T) {
	taskStore := MockTaskStore()
	userStore := MockUserStore()
	handler := NewHandler(userStore, taskStore)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	t.Run("Create Task", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Test Task",
			Description: "This is a test task",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
		token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
		req.Header.Set("Authorization", token)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)

		// Adicione asserções adicionais se necessário para a criação de tarefa
	})

	t.Run("List Tasks", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
		token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
		req.Header.Set("Authorization", token)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		// Adicione asserções adicionais se necessário para listagem de tarefas
	})

	t.Run("Get Task", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/tasks/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
		token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
		req.Header.Set("Authorization", token)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		// Adicione asserções adicionais se necessário para obtenção de tarefa
	})

	t.Run("Update Task", func(t *testing.T) {
		payload := types.TaskPayLoad{
			Title:       "Updated Task",
			Description: "This is an updated task",
		}
		body, _ := json.Marshal(payload)
		req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
		token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
		req.Header.Set("Authorization", token)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		// Adicione asserções adicionais se necessário para atualização de tarefa
	})

	t.Run("Delete Task", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
		token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
		req.Header.Set("Authorization", token)

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		// Adicione asserções adicionais se necessário para deleção de tarefa
	})

	//  t.Run("Update Task Status", func(t *testing.T) {
	// 	req, err := http.NewRequest("PUT", "/tasks/1/update", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	// Simulando autenticação adicionando token JWT ao cabeçalho da requisição
	// 	token, _ := auth.CreateJWT([]byte(configs.Envs.JWTSecret), 1)
	// 	req.Header.Set("Authorization", token)

	// 	rr := httptest.NewRecorder()
	// 	mux.ServeHTTP(rr, req)

	// 	assert.Equal(t, http.StatusOK, rr.Code)

	// 	// Adicione asserções adicionais se necessário para atualização de status de tarefa
	// })

}

// mock methods for the TaskStore interface
func (m *mockTaskStore) CreateTask(payload types.TaskPayLoad, userID int) error {
	taskID := len(m.tasks) + 1 // Simulando a criação de um ID único
	task := types.Task{
		ID:          taskID,
		Title:       payload.Title,
		Description: payload.Description,
		UserID:      userID,
		Status:      "ToDo", // Defina o status padrão conforme necessário
	}
	m.tasks[taskID] = task
	return nil
}

func (m *mockTaskStore) ListTasks(userID int) ([]*types.Task, error) {
	var userTasks []*types.Task
	for _, task := range m.tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, &task)
		}
	}
	return userTasks, nil
}

func (m *mockTaskStore) GetTaskByID(userID, taskID int) (*types.Task, error) {
	task, exists := m.tasks[taskID]
	if !exists {
		return nil, errors.New("task not found")
	}
	if task.UserID != userID {
		return nil, errors.New("user does not have access to this task")
	}
	return &task, nil
}

func (m *mockTaskStore) UpdateTask(updatedTask types.Task) error {
	_, exists := m.tasks[updatedTask.ID]
	if !exists {
		return errors.New("task not found")
	}
	m.tasks[updatedTask.ID] = updatedTask
	return nil
}

func (m *mockTaskStore) DeleteTask(userID, taskID int) error {
	_, exists := m.tasks[taskID]
	if !exists {
		return errors.New("task not found")
	}
	delete(m.tasks, taskID)
	return nil
}

// mock methods for the UserStore interface
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
