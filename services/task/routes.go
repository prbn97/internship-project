package task

import (
	"cmd/main.go/types"
	"cmd/main.go/utils"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
)

var (
	getTodoRegularExpression      = regexp.MustCompile(`^/tasks/([a-zA-Z0-9]+)$`)
	completeTodoRegularExpression = regexp.MustCompile(`^/tasks/([a-zA-Z0-9]+)/complete$`)
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
	serv.HandleFunc("POST /tasks", h.taskPOST)
	serv.HandleFunc("GET /tasks", h.taskLIST)
	serv.HandleFunc("GET /tasks/{id}", h.taskGET)
	serv.HandleFunc("PUT /tasks/{id}", h.taskPUT)
	serv.HandleFunc("DELETE /tasks/{id}", h.taskDELETE)
	serv.HandleFunc("PUT /tasks/{id}/complete", h.taskComplete)
}

func (h *Handler) taskPOST(w http.ResponseWriter, r *http.Request) {

	var task types.TaskPayLoad

	err := utils.ValidateFields(w, r, &task)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if task.Title == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("task title is required"))
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err = h.store.CreateTask(task)
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

	var updatedTask types.TaskPayLoad

	err = utils.ValidateFields(w, r, &updatedTask)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if updatedTask.Title == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("task title is required"))
		return
	}

	if updatedTask.Title != task.Title {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != task.Description {
		task.Description = updatedTask.Description
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
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("this task is already completed"))
		return
	}

	task.Completed = true
	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
