package handlers

import (
	"api/main.go/models"
	"api/main.go/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
)

var (
	listUser_RegularExpression   = regexp.MustCompile(`^\/users[\/]*$`)
	getUser_RegularExpression    = regexp.MustCompile(`^\/users\/([a-zA-Z0-9]+)$`)
	createUser_RegularExpression = regexp.MustCompile(`^\/users[\/]*$`)
)

type Datastore struct {
	m map[string]models.User
	*sync.RWMutex
}

type UserHandler struct {
	store *Datastore
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		store: &Datastore{
			m:       map[string]models.User{},
			RWMutex: &sync.RWMutex{},
		},
	}
}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")

	switch {
	case req.Method == http.MethodGet && listUser_RegularExpression.MatchString(req.URL.Path):
		h.List(res, req)
		return
	case req.Method == http.MethodGet && getUser_RegularExpression.MatchString(req.URL.Path):
		h.Get(res, req)
		return
	case req.Method == http.MethodPost && createUser_RegularExpression.MatchString(req.URL.Path):
		h.Create(res, req)
		return
	case req.Method == http.MethodPut && getUser_RegularExpression.MatchString(req.URL.Path):
		h.Update(res, req)
		return
	case req.Method == http.MethodDelete && getUser_RegularExpression.MatchString(req.URL.Path):
		h.Delete(res, req)
		return
	default:
		utils.NotFound(res, req)
	}

}

func (h *UserHandler) List(res http.ResponseWriter, req *http.Request) {
	h.store.RLock()
	defer h.store.RUnlock()

	users := make([]models.User, 0, len(h.store.m))
	for _, user := range h.store.m {
		users = append(users, user)
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *UserHandler) Get(res http.ResponseWriter, req *http.Request) {
	matches := getUser_RegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}
	h.store.RLock()
	user, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *UserHandler) Create(res http.ResponseWriter, req *http.Request) {

	u := models.User{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		utils.BadRequest(res, req)
		return
	}

	newID, err := GenerateID(20)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	u.ID = newID

	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)

}

func (h *UserHandler) Update(res http.ResponseWriter, req *http.Request) {
	matches := getUser_RegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}

	userID := matches[1]
	h.store.Lock()
	u, ok := h.store.m[userID]
	h.store.Unlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(req.Body).Decode(&updatedUser); err != nil {
		utils.BadRequest(res, req)
		return
	}
	if updatedUser.Name != "" {
		u.Name = updatedUser.Name
	}
	h.store.m[userID] = u

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *UserHandler) Delete(res http.ResponseWriter, req *http.Request) {
	matches := getUser_RegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		utils.NotFound(res, req)
		return
	}

	userID := matches[1]
	h.store.Lock()
	_, ok := h.store.m[userID]
	defer h.store.Unlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	delete(h.store.m, userID)
	res.WriteHeader(http.StatusOK)
}
