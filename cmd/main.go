package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/prbn97/internship-project/services/task"
)

const port = 8080

func main() {

	// build storage
	dbPath := filepath.Join("db", "tasksStore.json")
	tasksDB, err := task.NewStore(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize task store: %v", err)
	}

	// build API routes
	handler := http.NewServeMux()
	taskHandler := task.NewHandler(tasksDB)
	taskHandler.RegisterRoutes(handler)

	// servs info
	log.Printf("API running at http://localhost:%d/", port)
	log.Print("Listening...")

	// start the server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
