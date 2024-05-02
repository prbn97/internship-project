package main

import (
	"fmt"
	"net/http"
	"os"
)

const defaultPort = "3000"

func main() {
	// http.HandleFunc("/hello", hello)
	port := getPort()
	http.HandleFunc("/todo", TodoEntryPoint)
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
	fmt.Println("API running at http://localhost:" + port + "/todo")
}
