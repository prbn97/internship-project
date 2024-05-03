package main

import (
	"api/main.go/api/handlers"
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := getPort()
	http.HandleFunc("/", handlers.TodoEntryPoint)
	printServerInfo(port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}

func printServerInfo(port string) {
	fmt.Println("API running at http://localhost:" + port + "/")
}
