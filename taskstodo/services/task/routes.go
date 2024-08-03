package task

import (
	"fmt"
	"net/http"

	types "github.com/prbn97/internship-project/services/models"
	"github.com/prbn97/internship-project/services/utils"
)

type Handler struct {
	store types.TaskStore
}

func NewHandler(store types.TaskStore) *Handler {
	return &Handler{
		store: store,
	}
}

// go 1.22 feature "Method based routing"
func (h *Handler) RegisterRoutes(serv *http.ServeMux) {
	serv.HandleFunc("POST /tasks", h.handlePOST)
	serv.HandleFunc("GET /tasks", h.handleLIST)
	serv.HandleFunc("GET /tasks/{id}", h.handleGET)
	serv.HandleFunc("PUT /tasks/{id}", h.handlePUT)
	serv.HandleFunc("DELETE /tasks/{id}", h.handleDELETE)
	serv.HandleFunc("PUT /tasks/{id}/update", h.handleSTATUS)
}

func validateID(r *http.Request) (string, error) {
	// go 1.22 feature PathValue returns the value "path wildcard"
	id := r.PathValue("id")
	if len(id) != 20 {
		return "", fmt.Errorf("invalid id")
	}
	return id, nil
}

func (h *Handler) handlePOST(w http.ResponseWriter, r *http.Request) {

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

	err = h.store.CreateTask(task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, task)
}

func (h *Handler) handleLIST(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.store.ListTasks()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, tasks)
}

func (h *Handler) handleGET(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.store.GetTaskByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) handlePUT(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var updatedTask types.TaskPayLoad
	err = utils.ValidateFields(w, r, &updatedTask)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	task, err := h.store.GetTaskByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	if updatedTask.Title == "" && updatedTask.Description == task.Description {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("at least a new task description field must be provided"))
		return
	}
	if updatedTask.Title == "" {
		updatedTask.Title = task.Title
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

func (h *Handler) handleDELETE(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.store.DeleteTask(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}

func (h *Handler) handleSTATUS(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	task, err := h.store.GetTaskByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("not found id"))
		return
	}

	statuses := []string{"ToDo", "Doing", "Done"}
	var nextStatus string

	for i, status := range statuses {
		if task.Status == status {
			nextStatus = statuses[(i+1)%len(statuses)]
			break
		}
	}

	task.Status = nextStatus

	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
