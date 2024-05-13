package main

import (
	"api/main.go/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	setupHandlers(mux)

	fmt.Println("API running at http://localhost:8080")
	fmt.Println("Listening...")

	http.ListenAndServe("localhost:8080", mux)
}

func setupHandlers(mux *http.ServeMux) {

	userH := handlers.NewUserHandler()
	todoH := handlers.NewTodoHandler()

	mux.Handle("/users/", userH) // /users/{id}
	mux.Handle("/users", userH)  // /users

	mux.Handle("/todos/", todoH) // /todos/{id}
	mux.Handle("/todos", todoH)  // /todos
}
