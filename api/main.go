//make sure to break your code in multiple functions and when it makes sense, break it in different packages and files

package main

import (
	"api/main.go/internal/todo"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const name = "Paulo \"The King\""

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/todo", TodoEntryPoint)

	if os.Getenv("PORT") == "" {
		fmt.Printf("missing PORT\n")
		os.Exit(1)
	}
	fmt.Println("listening...")

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	message := fmt.Sprintf("Hello %s\n", name)
	res.WriteHeader(http.StatusOK) //Status code 200
	res.Write([]byte(message))
}

func TodoEntryPoint(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		CreateTodo(res, req)
		return
	} else if req.Method == "GET" {
		hello(res, req)
	}

}

func CreateTodo(res http.ResponseWriter, req *http.Request) {

	var newTodo todo.Todo
	err := json.NewDecoder(req.Body).Decode(&newTodo)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid todo item"))
		return
	}
	newTodo.ID = "1" //implementar um contador
	fmt.Println(newTodo)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(newTodo)

}
