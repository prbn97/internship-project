package tasks

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/prbn97/internship-project/services/auth"
	types "github.com/prbn97/internship-project/services/models"
	"github.com/prbn97/internship-project/services/utils"
)

type Handler struct {
	userStore types.UserStore
	store     types.TaskStore
}

func NewHandler(userStore types.UserStore, store types.TaskStore) *Handler {
	return &Handler{
		userStore: userStore,
		store:     store,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /tasks", auth.WithJWTAuth(h.handleCreateTask, h.userStore))
	mux.HandleFunc("GET /tasks", auth.WithJWTAuth(h.handleListTasks, h.userStore))

	mux.HandleFunc("GET /tasks/{id}", auth.WithJWTAuth(h.handleGetTask, h.userStore))
	mux.HandleFunc("PUT /tasks/{id}", auth.WithJWTAuth(h.handleUpdateTask, h.userStore))
	mux.HandleFunc("DELETE /tasks/{id}", auth.WithJWTAuth(h.handlDeleteTask, h.userStore))

	mux.HandleFunc("PUT /tasks/{id}/update", auth.WithJWTAuth(h.handleUpdateStatus, h.userStore))

}

// Get user and task IDs
func validateID(r *http.Request) (int, int, error) {

	userID := auth.GetUserIDFromContext(r.Context())
	// go 1.22 feature, returns the value "path wildcard {id}"
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, 0, fmt.Errorf("%s is a invalid id", idStr)
	}

	return userID, id, nil
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, r *http.Request) {

	// Get userID and payload in JSON
	userID := auth.GetUserIDFromContext(r.Context())

	var payload types.TaskPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(" invalid payload: %v", errors))
		return
	}

	// Create task
	err := h.store.CreateTask(payload, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleListTasks(w http.ResponseWriter, r *http.Request) {

	userID := auth.GetUserIDFromContext(r.Context())
	tasks, err := h.store.ListTasks(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)
}

func (h *Handler) handleGetTask(w http.ResponseWriter, r *http.Request) {

	// Get user and task IDs
	userID, taskID, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.store.GetTaskByID(userID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) handleUpdateTask(w http.ResponseWriter, r *http.Request) {

	userID, taskID, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var payload types.TaskPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(" invalid payload: %v", errors))
		return
	}

	task, err := h.store.GetTaskByID(userID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if payload.Title == "" && payload.Description == task.Description {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("new task description field must be provided"))
		return
	}
	if payload.Title == "" {
		payload.Title = task.Title
	}
	if payload.Title != task.Title {
		task.Title = payload.Title
	}
	if payload.Description != task.Description {
		task.Description = payload.Description
	}

	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)

}

func (h *Handler) handlDeleteTask(w http.ResponseWriter, r *http.Request) {

	// Get user and task IDs
	userID, taskID, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.store.GetTaskByID(userID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	err = h.store.DeleteTask(userID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)

}

func (h *Handler) handleUpdateStatus(w http.ResponseWriter, r *http.Request) {
	// Get user and task IDs
	userID, taskID, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	task, err := h.store.GetTaskByID(userID, taskID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	// Change for the next string status and update
	taskStatus := []string{"ToDo", "Doing", "Done"}
	var nextStatus string

	for i, status := range taskStatus {
		if task.Status == status {
			nextStatus = taskStatus[(i+1)%len(taskStatus)]
			break
		}
	}

	// Log next status
	fmt.Printf("Next task status: %s\n", nextStatus)

	task.Status = nextStatus
	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
