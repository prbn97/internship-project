package handlers

import (
	"api/main.go/models"
	"api/main.go/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

// expressões para manipular os endpoints de ToDo
var (
	listTodoRegularExpression   = regexp.MustCompile(`^/todos[\/]*$`)
	getTodoRegularExpression    = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)$`)
	createTodoRegularExpression = regexp.MustCompile(`^/todos[\/]*$`)
)

// TodoDatastore é uma estrutura para armazenar os itens de ToDo
type TodoDatastore struct {
	m map[string]models.Todo
	*sync.RWMutex
}

// TodoHandler é o manipulador para os endpoints relacionados a ToDo
type TodoHandler struct {
	store *TodoDatastore
}

// NewTodoHandler cria um novo manipulador para itens de ToDo
func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		store: &TodoDatastore{
			m:       make(map[string]models.Todo),
			RWMutex: &sync.RWMutex{},
		},
	}
}

// implementa o método ServeHTTP da interface http.Handler
func (h *TodoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	switch {
	case req.Method == http.MethodGet && listTodoRegularExpression.MatchString(req.URL.Path):
		h.List(res, req)
		return
	case req.Method == http.MethodGet && getTodoRegularExpression.MatchString(req.URL.Path):
		h.Get(res, req)
		return
	case req.Method == http.MethodPost && createTodoRegularExpression.MatchString(req.URL.Path):
		h.Create(res, req)
		return
	case req.Method == http.MethodPut && getTodoRegularExpression.MatchString(req.URL.Path):
		h.Update(res, req)
		return
	case req.Method == http.MethodDelete && getTodoRegularExpression.MatchString(req.URL.Path):
		h.Delete(res, req)
	default:
		utils.NotFound(res, req)
	}
}

func (h *TodoHandler) List(res http.ResponseWriter, req *http.Request) {
	h.store.RLock()
	defer h.store.RUnlock()

	todos := make([]models.Todo, 0, len(h.store.m))
	for _, todoItem := range h.store.m {
		todos = append(todos, todoItem) // ponteiro no todoItem
	}

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *TodoHandler) Get(res http.ResponseWriter, req *http.Request) {
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}
	h.store.RLock()
	todo, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {
	// decodes the JSON data from the request body into a user struct.
	// This POST assumes that the request contains JSON data representing a user.
	h.store.Lock()
	defer h.store.Unlock()

	u := models.Todo{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		utils.BadRequest(res, req)
		return
	}

	newID, err := GenerateID(20)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	u.ID = newID

	h.store.m[u.ID] = u

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)

}

func (h *TodoHandler) Update(res http.ResponseWriter, req *http.Request) {
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}

	todoID := matches[1]
	h.store.Lock()
	todoItem, ok := h.store.m[todoID]
	defer h.store.Unlock()

	if !ok {
		utils.NotFound(res, req)
		return
	}

	// Decodificar o corpo da solicitação para obter os novos dados do todo item
	var updatedTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
		utils.BadRequest(res, req)
		return
	}
	// Atualizar os campos do todo item, se fornecidos
	if updatedTodo.Title != "" {
		todoItem.Title = updatedTodo.Title
	}
	if updatedTodo.Description != "" {
		todoItem.Description = updatedTodo.Description
	}
	if updatedTodo.Completed != todoItem.Completed {
		todoItem.Completed = updatedTodo.Completed
	}
	h.store.m[todoID] = todoItem

	jsonBytes, err := json.Marshal(todoItem)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *TodoHandler) Delete(res http.ResponseWriter, req *http.Request) {
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}

	todoID := matches[1]
	h.store.Lock()
	_, ok := h.store.m[todoID]
	defer h.store.Unlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	delete(h.store.m, todoID)
	res.WriteHeader(http.StatusOK)
}
