//make sure to break your code in multiple functions and when it makes sense, break it in different packages and files

package main

import (
	"api/main.go/internal/todo"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "3000"

func main() {
	// http.HandleFunc("/hello", hello)
	port := getPort()
	http.HandleFunc("/todo", TodoEntryPoint)
	printServerInfo(port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}

func printServerInfo(port string) {
	fmt.Println("API running at http://localhost:" + port + "/todo")
}

func TodoEntryPoint(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		CreateTodo(res, req)
		return
	} else if req.Method == "GET" {
		hello(res, req)
	}

}

// Em api/todo_handler.go
func CreateTodo(res http.ResponseWriter, req *http.Request) {

	var newTodo todo.Todo
	err := json.NewDecoder(req.Body).Decode(&newTodo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid todo item"))
		return
	}

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

const name = "Paulo \"The King\""

func hello(res http.ResponseWriter, req *http.Request) {
	message := fmt.Sprintf("Hello %s\n", name)
	res.WriteHeader(http.StatusOK) //Status code 200
	res.Write([]byte(message))
}
