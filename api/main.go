package main

import (
	"api/main.go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

var (
	listUser_RegularExpression   = regexp.MustCompile(`^\/users[\/]*$`)
	getUser_RegularExpression    = regexp.MustCompile(`^\/users\/(\d+)$`)
	createUser_RegularExpression = regexp.MustCompile(`^\/users[\/]*$`)
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type datastore struct {
	m map[string]user
	*sync.RWMutex
}
type userHandler struct {
	store *datastore
}

func (h *userHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
	default:
		utils.NotFound(res, req)
	}

}

func (h *userHandler) List(res http.ResponseWriter, req *http.Request) {
	users := make([]user, 0, len(h.store.m))
	h.store.RLock()
	for _, user := range h.store.m {
		users = append(users, user)
	}

	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *userHandler) Get(res http.ResponseWriter, req *http.Request) {
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

func (h *userHandler) Create(res http.ResponseWriter, req *http.Request) {
	// decodes the JSON data from the request body into a user struct.
	// This POST assumes that the request contains JSON data representing a user.

	u := user{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		utils.BadRequest(res, req)
		return
	}

	newID := strconv.Itoa(len(h.store.m) + 1) // Generate a new ID as an integer
	u.ID = newID

	h.store.Lock()
	h.store.m[u.ID] = u // adds the new user to the datastore
	h.store.Unlock()

	res.Header().Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)

}

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{m: map[string]user{}, RWMutex: &sync.RWMutex{}},
	}
	fmt.Println("API running at http://localhost:8080/users")
	fmt.Println("listening...")

	mux.Handle("/users/", userH) // /users/{id}
	mux.Handle("/users", userH)  // /users
	http.ListenAndServe("localhost:8080", mux)

}
