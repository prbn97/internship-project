package users

import "net/http"

type Handler struct {
	// the Handler can take any dependence
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(serv *http.ServeMux) {

}
