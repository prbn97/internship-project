package handlers

import (
	"api/main.go/models"
	"api/main.go/utils"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
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
	// decodes the JSON data from the request body into a user struct.
	// This POST assumes that the request contains JSON data representing a user.

	u := models.User{}
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		utils.BadRequest(res, req)
		return
	}

	newID, err := generateID(20) // Generate a new ID as an integer
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}

	u.ID = newID

	h.store.Lock()
	h.store.m[u.ID] = u // adds the new user to the datastore
	h.store.Unlock()

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerError(res, req)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(jsonBytes)

}
func generateID(length int) (string, error) {
	// Calcula o número de bytes necessário para gerar o ID
	numBytes := length / 2
	if length%2 != 0 {
		numBytes++
	}

	// Gera bytes aleatórios usando crypto/rand
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Codifica os bytes aleatórios em uma string hexadecimal
	id := hex.EncodeToString(randomBytes)

	// Ajusta o tamanho do ID se necessário
	if len(id) > length {
		id = id[:length]
	} else if len(id) < length {
		// Se o ID gerado for menor que o tamanho especificado,
		// preenche o restante com caracteres '0'
		id += strings.Repeat("0", length-len(id))
	}

	return id, nil
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
	// Decodificar o corpo da solicitação para obter os novos dados do usuário
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
	h.store.Unlock()
	if !ok {
		utils.NotFound(res, req)
		return
	}

	delete(h.store.m, userID)
	res.WriteHeader(http.StatusOK)
}
