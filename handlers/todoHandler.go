package handlers

import (
	"api/main.go/models"
	"api/main.go/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listTodoRegularExpression   = regexp.MustCompile(`^/todos[\/]*$`)
	getTodoRegularExpression    = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)$`)
	createTodoRegularExpression = regexp.MustCompile(`^/todos[\/]*$`)
)

type TodoDatastore struct {
	m map[string]models.Todo
	*sync.RWMutex
}

type TodoHandler struct {
	store *TodoDatastore
}

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		store: &TodoDatastore{
			m:       make(map[string]models.Todo),
			RWMutex: &sync.RWMutex{},
		},
	}
}

func (h *TodoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")

	switch {
	case req.Method == http.MethodPost && createTodoRegularExpression.MatchString(req.URL.Path):
		h.Create(res, req)
		return

	case req.Method == http.MethodGet && listTodoRegularExpression.MatchString(req.URL.Path):
		h.List(res, req)
		return

	case req.Method == http.MethodGet: // && getTodoRegularExpression.MatchString(req.URL.Path):
		h.Get(res, req)
		return

	case req.Method == http.MethodPut: // && getTodoRegularExpression.MatchString(req.URL.Path)
		h.Update(res, req)
		return

	case req.Method == http.MethodDelete: // && getTodoRegularExpression.MatchString(req.URL.Path):
		h.Delete(res, req)
		return

	default:
		utils.MethodNotAllowed(res, req)
	}
}

func (h *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	h.store.Lock()
	defer h.store.Unlock()

	var todo models.Todo
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields() // this will cause decoder to return an error if any unknown field is encountered
	if err := decoder.Decode(&todo); err != nil {
		utils.BadRequest(res, req, "invalid json")
		return
	}
	// Ensure the ID is not provided by the user
	if todo.ID != "" {
		utils.BadRequest(res, req, "ID cannot be provided")
		return
	}

	// Ensure the title is provided by the user
	if todo.Title == "" {
		utils.BadRequest(res, req, "title is required")
		return
	}

	newID, err := GenerateID(20)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	todo.ID = newID

	h.store.m[todo.ID] = todo

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusCreated)
	res.Write(jsonBytes)
}

func (h *TodoHandler) List(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	h.store.RLock()
	defer h.store.RUnlock()

	todos := make([]models.Todo, 0, len(h.store.m))
	for _, todoItem := range h.store.m {
		todos = append(todos, todoItem)
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
	res.Header().Set("content-type", "application/json")
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.NotFound(res, req, "invalid ID")
		return
	}

	h.store.RLock()
	defer h.store.RUnlock()

	todo, ok := h.store.m[matches[1]]
	if !ok {
		utils.NotFound(res, req, "ID not found")
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

func (h *TodoHandler) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.NotFound(res, req, "invalid ID")
		return
	}

	todoID := matches[1]

	h.store.Lock()
	defer h.store.Unlock()

	todoItem, ok := h.store.m[todoID]
	if !ok {
		utils.NotFound(res, req, "ID not found")
		return
	}

	var updatedTodo models.Todo
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields() // this will cause decoder to return an error if any unknown field is encountered
	if err := decoder.Decode(&updatedTodo); err != nil {
		utils.BadRequest(res, req, "invalid json")
		return
	}
	// Ensure the ID in the payload is not being updated
	if updatedTodo.ID != "" && updatedTodo.ID != todoID {
		utils.BadRequest(res, req, "ID cannot be updated")
		return
	}
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
	res.Header().Set("content-type", "application/json")
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.NotFound(res, req, "invalid ID")
		return
	}

	todoID := matches[1]

	h.store.Lock()
	defer h.store.Unlock()

	todoItem, ok := h.store.m[todoID]
	if !ok {
		utils.NotFound(res, req, "ID not found")
		return
	}
	jsonBytes, err := json.Marshal(todoItem)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	delete(h.store.m, todoID)
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}
