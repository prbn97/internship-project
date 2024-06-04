package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GenerateID(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length")
	}
	numBytes := length / 2
	if length%2 != 0 {
		numBytes++
	}
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	id := hex.EncodeToString(randomBytes)

	if len(id) > length {
		id = id[:length]
	} else if len(id) < length {
		id += strings.Repeat("0", length-len(id))
	}

	return id, nil
}

// func TestGenerateID(t *testing.T) {
// 	id, err := GenerateID(20)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 20, len(id))

// 	_, err = GenerateID(-25)
// 	assert.Error(t, err)
// }
