package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/prbn97/internship-project/services/task"
)

func main() {

	loadEnvFiles()

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

func loadEnvFiles() {

	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found. Trying to load env.example")

		if err := godotenv.Load(".env.example"); err != nil {
			log.Println("No .env or env.example files found")
		} else {
			log.Println("env.example file loaded")
		}

	} else {
		log.Println(".env file loaded")
	}
}
