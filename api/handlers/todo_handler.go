package handlers

import (
	todo "api/main.go/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var todoList []todo.Todo

func TodoEntryPoint(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		CreateTodoHandler(res, req)
		return
	} else if req.Method == "GET" {
		GetAllTodosHandler(res, req)
	}

	// GET /todo/{id} (retrieve a single TODO item by ID)
	// PUT /todo/{id} (update a TODO item by ID)
	// DELETE /todo/{id} (delete a TODO item by ID)

}

func CreateTodoHandler(res http.ResponseWriter, req *http.Request) {

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

func GetAllTodosHandler(res http.ResponseWriter, req *http.Request) {
	// Carrega a lista de todo items do arquivo JSON
	todoList, err := todo.LoadListFromFile("todos.json")
	if err != nil {
		fmt.Println("Erro ao carregar a lista de todo items:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todoList)
}
