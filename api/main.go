package main

import (
	"api/main.go/handlers"
	"fmt"
	"net/http"
)

// create todo_handler

func main() {
	mux := http.NewServeMux()
	userH := handlers.NewUserHandler()

	fmt.Println("API running at http://localhost:8080/users")
	fmt.Println("listening...")

	mux.Handle("/users/", userH) // /users/{id}
	mux.Handle("/users", userH)  // /users
	http.ListenAndServe("localhost:8080", mux)

}
