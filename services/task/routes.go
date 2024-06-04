package task

import (
	"api/main.go/types"
	"api/main.go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var (
	getTodoRegularExpression      = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)$`)
	completeTodoRegularExpression = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)/complete$`)
)

type TasksStore struct {
	m map[string]types.Task
}

type Handler struct {
	store *TasksStore
}

func NewHandler() *Handler {
	return &Handler{
		store: &TasksStore{
			m: make(map[string]types.Task)},
	}
}

func (h *Handler) RegisterRoutes(serv *http.ServeMux) {

	serv.HandleFunc("POST /todos", h.Create)
	serv.HandleFunc("GET /todos/", h.List)
	serv.HandleFunc("GET /todos/{id}", h.Get)
	serv.HandleFunc("PUT /todos/{id}", h.Update)
	serv.HandleFunc("DELETE /todos/{id}", h.Delete)
	serv.HandleFunc("PUT /todos/{id}/complete", h.MarkComplete)

}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var todo types.Task
	decoder := json.NewDecoder(r.Body)
	// this will cause decoder to return an error if any unknown field is encountered
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&todo); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid json"))
		return
	}
	// Ensure the ID is not provided by the user
	if todo.ID != "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("dont add id"))
		return
	}

	// Ensure the title is provided by the user
	if todo.Title == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("title cant be empty"))
		return
	}

	newID, err := utils.GenerateID(20)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	todo.ID = newID

	h.store.m[todo.ID] = todo

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	todos := make([]types.Task, 0, len(h.store.m))
	for _, todoItem := range h.store.m {
		todos = append(todos, todoItem)
	}

	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	matches := getTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	todo, ok := h.store.m[matches[1]]
	if !ok {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	jsonBytes, err := json.Marshal(todo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	matches := getTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	todoID := matches[1]

	todoItem, ok := h.store.m[todoID]
	if !ok {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	var updatedTodo types.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // this will cause decoder to return an error if any unknown field is encountered
	if err := decoder.Decode(&updatedTodo); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid json"))
		return
	}
	// Ensure the ID in the payload is not being updated
	if updatedTodo.ID != "" && updatedTodo.ID != todoID {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cant change id"))
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
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) Delete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	matches := getTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	todoID := matches[1]

	todoItem, ok := h.store.m[todoID]
	if !ok {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}
	jsonBytes, err := json.Marshal(todoItem)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	delete(h.store.m, todoID)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *Handler) MarkComplete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	matches := completeTodoRegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	todoID := matches[1]

	todoItem, ok := h.store.m[todoID]
	if !ok {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	// mark todo task as true
	if todoItem.Completed {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("this task completed"))
		return
	}
	todoItem.Completed = true
	h.store.m[todoID] = todoItem

	jsonBytes, err := json.Marshal(todoItem)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
