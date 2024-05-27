package main

import (
	"api/main.go/handlers"
	"log"
	"net/http"
)

func main() {
	Port := ":8080"
	ServeMux := http.NewServeMux()
	loadRoutes(ServeMux)

	server := http.Server{
		Addr:    Port,
		Handler: ServeMux,
	}

	log.Printf("API running at http://localhost: %s", Port)
	log.Print("Listening...")
	server.ListenAndServe()
}

func loadRoutes(router *http.ServeMux) {
	todoHandler := handlers.NewTodoHandler()

	router.HandleFunc("POST /todos", todoHandler.Create)
	router.HandleFunc("GET /todos/", todoHandler.List)
	router.HandleFunc("GET /todos/{id}", todoHandler.Get)
	router.HandleFunc("PUT /todos/{id}", todoHandler.Update)
	router.HandleFunc("DELETE /todos/{id}", todoHandler.Delete)
}
