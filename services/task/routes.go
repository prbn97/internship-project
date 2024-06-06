package task

import (
	"cmd/main.go/types"
	"cmd/main.go/utils"
	"fmt"
	"net/http"
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
	serv.HandleFunc("POST /tasks", h.taskPOST)
	serv.HandleFunc("GET /tasks", h.taskLIST)
	serv.HandleFunc("GET /tasks/{id}", h.taskGET)
	serv.HandleFunc("PUT /tasks/{id}", h.taskPUT)
	serv.HandleFunc("DELETE /tasks/{id}", h.taskDELETE)
	serv.HandleFunc("PUT /tasks/{id}/complete", h.taskComplete)
	serv.HandleFunc("PUT /tasks/{id}/incomplete", h.taskIncomplete)
}

func validateID(r *http.Request) (string, error) {
	// go 1.22 feature PathValue returns the value "path wildcard"
	id := r.PathValue("id")
	if len(id) != 20 {
		return "", fmt.Errorf("invalid id")
	}
	return id, nil
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

func (h *Handler) taskPUT(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) taskDELETE(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) taskComplete(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) taskIncomplete(w http.ResponseWriter, r *http.Request) {
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

	if !task.Completed {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("this task is already incompleted"))
		return
	}

	task.Completed = false
	err = h.store.UpdateTask(*task)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, task)
}
