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

	log.Printf("API running at http://localhost%s", Port)
	log.Print("Listening...")
	server.ListenAndServe()
}

func loadRoutes(serv *http.ServeMux) {
	todoHandler := handlers.NewTodoHandler()

	serv.HandleFunc("POST /todos", todoHandler.Create)
	serv.HandleFunc("GET /todos/", todoHandler.List)
	serv.HandleFunc("GET /todos/{id}", todoHandler.Get)
	serv.HandleFunc("PUT /todos/{id}", todoHandler.Update)
	serv.HandleFunc("DELETE /todos/{id}", todoHandler.Delete)
	serv.HandleFunc("PUT /todos/{id}/complete", todoHandler.MarkComplete)
}
