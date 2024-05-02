package main

import (
	"api/main.go/internal/todo"
	"encoding/json"
	"fmt"
	"net/http"
)

func TodoEntryPoint(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		CreateTodoHandler(res, req)
		return
	} else if req.Method == "GET" {
		hello(res, req)
	}
	// make now!
	// GET /todo (retrieve all TODO items)
	// GET /todo/{id} (retrieve a single TODO item by ID)

	// make latter
	// PUT /todo/{id} (update a TODO item by ID)
	// DELETE /todo/{id} (delete a TODO item by ID)

}

func CreateTodoHandler(res http.ResponseWriter, req *http.Request) {

	var newTodo todo.Todo
	err := json.NewDecoder(req.Body).Decode(&newTodo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid todo item"))
		return
	}

	// make now!
	// TODO: Implement logic to generate ID (e.g., UUID)

	newTodo.ID = "1"
	// newTodo.ID = generateID()

	// TODO: Implement logic to store the new Todo item
	// yourTodoStorageFunction(newTodo)

	// Return success response with the created Todo
	fmt.Println(newTodo)
	res.WriteHeader(http.StatusCreated) //201 Created
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(newTodo)

}

// make this my GET EndPoint
func hello(res http.ResponseWriter, req *http.Request) {
	name := "Paulo \"The King\""
	message := fmt.Sprintf("Hello %s\n", name)
	res.WriteHeader(http.StatusOK) //Status code 200
	res.Write([]byte(message))
}
