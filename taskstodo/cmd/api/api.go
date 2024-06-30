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

// the api holds the handlers and the stores
// manege the api services from here
func (serv *APIserver) Run() error {
	router := http.NewServeMux()

	// learn about subrouter and how to use.
	// subrouter := http.NewServeMux()
	// subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	userStore := users.NewStore(serv.database)
	usersRoutes := users.NewHandler(userStore)
	usersRoutes.RegisterRoutes(router)

	tasksStore := tasks.NewStore(serv.database)
	tasksRoutes := tasks.NewHandler(tasksStore)
	tasksRoutes.RegisterRoutes(router)

	// Listen and Serve router
	log.Println("Listening on", serv.addres)
	return http.ListenAndServe(serv.addres, router)
}
