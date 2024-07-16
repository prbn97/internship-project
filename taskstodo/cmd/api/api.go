package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/prbn97/internship-project/services/tasks"
	"github.com/prbn97/internship-project/services/users"
)

type APIserver struct {
	address  string
	database *sql.DB
}

func NewAPIserver(port string, db *sql.DB) *APIserver {
	return &APIserver{
		address:  port,
		database: db,
	}
}

// the api holds the handlers and the storages
func (serve *APIserver) Run() error {
	router := http.NewServeMux()

	// learn how implement subrouter.
	// subrouter := http.NewServeMux()
	// subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1/", router))

	// users service
	userStore := users.NewStore(serve.database)
	usersRoutes := users.NewHandler(userStore)
	usersRoutes.RegisterRoutes(router)

	// tasks service
	tasksStore := tasks.NewStore(serve.database)
	tasksRoutes := tasks.NewHandler(userStore, tasksStore)
	tasksRoutes.RegisterRoutes(router)

	log.Println("Listening on", serve.address)
	return http.ListenAndServe(serve.address, router)
}
