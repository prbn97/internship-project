package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/prbn97/internship-project/services/tasks"
	"github.com/prbn97/internship-project/services/users"
)

type APIserver struct {
	addres   string
	database *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addres:   addr,
		database: db,
	}
}

func (serv *APIserver) Run() error {
	// initialeze a router and subrouter
	router := http.NewServeMux()
	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	// registre all routes
	usersService := users.NewHandler()
	usersService.RegisterRoutes(subrouter)

	tasksService := tasks.NewHandler()
	tasksService.RegisterRoutes(subrouter)
	// Listen and Serve router
	log.Println("Listening on", serv.addres)
	return http.ListenAndServe(serv.addres, router)
}
