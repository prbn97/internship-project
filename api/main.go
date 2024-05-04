package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var (
	createUser_RegularExpression = regexp.MustCompile(`^\/users[\/]*$`)   // /post/
	listUser_RegularExpression   = regexp.MustCompile(`^\/users[\/]*$`)   // /user/
	getUser_RegularExpression    = regexp.MustCompile(`^\/users\/(\d+)$`) // /user/{id}
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
		notFound(res, req)
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
		internalServerError(res, req)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *userHandler) Get(res http.ResponseWriter, req *http.Request) {
	matches := getUser_RegularExpression.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
		notFound(res, req)
		return
	}
	h.store.RLock()
	user, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		notFound(res, req)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		internalServerError(res, req)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)
}

func (h *userHandler) Create(res http.ResponseWriter, req *http.Request) {
	u := user{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		badRequest(res, req)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()

	res.Header().Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)

}

func badRequest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(`{"error": "bad request"}`))
}

func notFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"error": "not found"}`))
}

func internalServerError(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`{"error": "internal server error"}`))
}

func main() {
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string]user{
				"1": {ID: "171", Name: "bob"},
				"2": {ID: "210", Name: "karen"},
				"3": {ID: "343", Name: "jack"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	fmt.Println("API running at http://localhost:8080/users")
	fmt.Println("listening...")

	mux.Handle("/users/", userH) // /users/{id}
	mux.Handle("/users", userH)  // /users/{id}
	http.ListenAndServe("localhost:8080", mux)

}

// func getPort() string {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		return "8080"
// 	}
// 	return port
// }

// func printServerInfo(port string) {
// 	fmt.Println("API running at http://localhost:" + port + "/")
// }
