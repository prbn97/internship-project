package handlers

import (
	"api/main.go/models"
	"api/main.go/utils"
	"encoding/json"
	"fmt"
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
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req) // error id not valid?
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

	h.store.Lock()
	defer h.store.Unlock()

	u := models.Todo{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		fmt.Println(err)
		utils.BadRequest(res, req, "invalid json")
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
		utils.NotFound(res, req) // error id not valid?
		return
	}

	todoID := matches[1]
	h.store.Lock()
	todoItem, ok := h.store.m[todoID]
	defer h.store.Unlock()

	if !ok {
		utils.NotFound(res, req) // error id not found
		return
	}

	var updatedTodo models.Todo
	if err := json.NewDecoder(req.Body).Decode(&updatedTodo); err != nil {
		utils.BadRequest(res, req, err.Error()) // what say?
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
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req) // error id not valid?
		return
	}

	todoID := matches[1]
	h.store.Lock()
	_, ok := h.store.m[todoID]
	defer h.store.Unlock()
	if !ok {
		utils.NotFound(res, req) // error id not found
		return
	}

	delete(h.store.m, todoID)
	res.WriteHeader(http.StatusOK)
}