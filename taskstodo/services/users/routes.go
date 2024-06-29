package users

import (
	"net/http"

	"github.com/prbn97/internship-project/types"
	"github.com/prbn97/internship-project/utils"
)

type Handler struct {
	// the Handler can take any dependence
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(serv *http.ServeMux) {
	serv.HandleFunc("POST /users/login", h.handleLogin)
	serv.HandleFunc("POST /users/register", h.handleRegister)

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// recive payloas to get eamil, user  and password
	// parse to json, validate
	// check if the user exist

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exists
	// if [false] create a user

}
