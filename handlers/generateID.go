package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
)

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
