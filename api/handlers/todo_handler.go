package handlers

import (
	todo "api/main.go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var todoList []todo.Todo

func TodoEntryPoint(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/todo" {
		switch req.Method {
		case "POST":
			CreateTodoItem(res, req)
		case "GET":
			GetAllTodoItem(res, req)
		default:
			res.WriteHeader(http.StatusMethodNotAllowed)
			res.Write([]byte("error method not allowed"))
		}
	} else if strings.HasPrefix(req.URL.Path, "/todo/") && req.Method == "GET" {
		GetTodoItemByID(res, req)
	} else {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("error endpoint not found"))
	}
	// PUT /todo/{id} (update a TODO item by ID)
	// DELETE /todo/{id} (delete a TODO item by ID)

}

func CreateTodoItem(res http.ResponseWriter, req *http.Request) {

	var newTodo todo.Todo
	err := json.NewDecoder(req.Body).Decode(&newTodo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("error generating ToDo Item"))
		return
	}

	id, err := todo.GenerateID()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("error generating  ID"))
		return
	}
	newTodo.ID = id
	todoList = append(todoList, newTodo)

	err = todo.SaveListOnFile(todoList, "todos.json")
	if err != nil {
		fmt.Println("error saving list of ToDo items", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return success response with the created Todo
	fmt.Println(newTodo)
	res.WriteHeader(http.StatusCreated) //201 Created
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(newTodo)

}

func GetAllTodoItem(res http.ResponseWriter, req *http.Request) {
	// Carrega a lista de todo items do arquivo JSON
	todoList, err := todo.LoadListFromFile("todos.json")
	if err != nil {
		fmt.Println("error when loading the list of todo items:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todoList)
}

func GetTodoItemByID(res http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")

	if len(parts) < 3 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("error missing ID"))
		return
	}

	id := parts[2]

	todoList, err := todo.LoadListFromFile("todos.json")
	if err != nil {
		fmt.Println("error when loading the list of ToDo items:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var foundTodo *todo.Todo
	for _, todo := range todoList {
		if todo.ID == id {
			foundTodo = &todo
			break
		}
	}

	if foundTodo == nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("error ToDo Item not find "))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(foundTodo)
}
