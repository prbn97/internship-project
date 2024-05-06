package main

import (
	"api/main.go/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Configure os handlers
	setupHandlers(mux)

	fmt.Println("API running at http://localhost:8080")
	fmt.Println("Listening...")

	// Inicie o servidor HTTP
	http.ListenAndServe("localhost:8080", mux)
}

func setupHandlers(mux *http.ServeMux) {
	// Inicialize os handlers de usuário e de ToDo
	userH := handlers.NewUserHandler()
	todoH := handlers.NewTodoHandler()

	// Registre os handlers para seus respectivos endpoints
	mux.Handle("/users/", userH) // /users/{id}
	mux.Handle("/users", userH)  // /users

	mux.Handle("/todos/", todoH) // /todos/{id}
	mux.Handle("/todos", todoH)  // /todos
}
