package main

import (
	"api/main.go/services/task"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	Port := ":8080"
	dbPath := filepath.Join("db", "tasks.json")
	store, err := task.NewStore(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize task store: %v", err)
	}

	handler := task.NewHandler(store)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	server := http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Printf("API running at http://localhost%s", Port)
	log.Print("Listening...")
	server.ListenAndServe()
}
