package main

import (
	"cmd/main.go/services/task"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// environment config
	Port := ":8080"
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    Port,
		Handler: mux,
	}

	// build API storage service
	dbPath := filepath.Join("db", "tasksStore.json")
	tasksDB, err := task.NewStore(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize task store: %v", err)
	}

	// build API routes services
	hTask := task.NewHandler(tasksDB)
	hTask.RegisterRoutes(mux)

	// servs info
	log.Printf("API running at http://localhost%s/", Port)
	log.Print("Listening...")

	server.ListenAndServe()

}
