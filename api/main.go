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

// curl --request POST \
//   --data '{"title":"task title", "description": "task description"}'\
//   http://localhost:8080/todo

// curl --request POST \
//   --data '{"title":"task title", "description": "remvover esse"}' \
//   http://localhost:8080/todo

// curl --request GET \
//   http://localhost:8080/todo/d75e8f8c649d2fc1879c2a645dd582d1
