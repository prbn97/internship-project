package main

import (
	"log"

	"github.com/prbn97/internship-project/cmd/api"
)

// the entry point for the API
func main() {
	// create a API serv instance
	server := api.NewAPIserver(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
