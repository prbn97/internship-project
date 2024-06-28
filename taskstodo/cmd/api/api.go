package api

import (
	"database/sql"
	"net/http"

	"github.com/prbn97/internship-project/services/tasks"
	"github.com/prbn97/internship-project/services/users"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, database *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   database,
	}
}

func (serv *APIserver) Run() error {
	// initialeze a router
	router := http.NewServeMux()

	// registre handlers
	users := users.NewHandler()
	users.RegisterRoutes(router)
	tasks := tasks.NewHandler()
	tasks.RegisterRoutes(router)

	// registres all dependencies
	return nil
}
