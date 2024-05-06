package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func GenerateID(length int) (string, error) {
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
