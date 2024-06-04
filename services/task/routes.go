package task

import (
	"api/main.go/types"
	"api/main.go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
)

var (
	getTodoRegularExpression      = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)$`)
	completeTodoRegularExpression = regexp.MustCompile(`^/todos/([a-zA-Z0-9]+)/complete$`)
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{
		store: store,
	}
}
func (h *Handler) RegisterRoutes(serv *http.ServeMux) {
	serv.HandleFunc("POST /todos", h.taskPOST)
	serv.HandleFunc("GET /todos", h.taskLIST)
	serv.HandleFunc("GET /todos/{id}", h.taskGET)
	serv.HandleFunc("PUT /todos/{id}", h.taskPUT)
	serv.HandleFunc("DELETE /todos/{id}", h.taskDELETE)
	serv.HandleFunc("PUT /todos/{id}/complete", h.taskComplete)
}

func (h *Handler) taskPOST(w http.ResponseWriter, r *http.Request) {
	var task types.TaskPayLoad
	if err := utils.ParseJSON(r, &task); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.store.CreateTask(task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, task)
}

func (h *Handler) taskLIST(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.store.ListTasks()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)
}

func (h *Handler) taskGET(w http.ResponseWriter, r *http.Request) {
	matches := getTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	task, err := h.store.GetTaskByID(matches[1])
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) taskPUT(w http.ResponseWriter, r *http.Request) {

	matches := getTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	task, err := h.store.GetTaskByID(matches[1])
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	var updatedTask types.Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&updatedTask); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if updatedTask.ID != "" && updatedTask.ID != matches[1] {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("cant change id"))
		return
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != task.Description {
		task.Description = updatedTask.Description
	}
	if updatedTask.Completed != task.Completed {
		task.Completed = updatedTask.Completed
	}

	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) taskDELETE(w http.ResponseWriter, r *http.Request) {

	matches := getTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	task, err := h.store.DeleteTask(matches[1])
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) taskComplete(w http.ResponseWriter, r *http.Request) {

	matches := completeTodoRegularExpression.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 || len(matches[1]) != 20 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	task, err := h.store.GetTaskByID(matches[1])
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	if task.Completed {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("this task completed"))
		return
	}

	task.Completed = true
	err = h.store.UpdateTask(types.Task{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
