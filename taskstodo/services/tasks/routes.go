package tasks

import (
	"net/http"

	"github.com/prbn97/internship-project/types"
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

}
