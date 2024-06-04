package main

import (
	"api/main.go/services/task"
	"log"
	"net/http"
)

func main() {
	Port := ":8080"
	ServeMux := http.NewServeMux()
	taskHandler := task.NewHandler()
	taskHandler.RegisterRoutes(ServeMux)

	server := http.Server{
		Addr:    Port,
		Handler: ServeMux,
	}

	log.Printf("API running at http://localhost%s", Port)
	log.Print("Listening...")
	server.ListenAndServe()
}
