package models

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GenerateID() (string, error) {
	idBytes := make([]byte, 16)

	_, err := rand.Read(idBytes)
	if err != nil {
		return "", err
	}
	id := hex.EncodeToString(idBytes)

	return id, nil
}

func SaveListOnFile(todoList []Todo, filename string) error {
	jsonData, err := json.Marshal(todoList)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadListFromFile(filename string) ([]Todo, error) {
	var todoList []Todo

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decodifica o conteúdo JSON para a lista de todo items
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}
