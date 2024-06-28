package users

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(serv *http.ServeMux) {

}
