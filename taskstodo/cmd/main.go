package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/prbn97/internship-project/services/task"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

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

	// wrap the handler with the CORS middleware
	corsHandler := enableCORS(handler)

	// servs info
	log.Printf("API running at http://localhost:%s/", os.Getenv("PORT"))
	log.Print("Listening...")

	// start the server
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), corsHandler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
