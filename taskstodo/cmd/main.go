package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/prbn97/internship-project/cmd/api"
	"github.com/prbn97/internship-project/config"
	"github.com/prbn97/internship-project/db"
)

// the entry point for the API
func main() {
	// create the database
	db, err := db.NewMySQLstorage(mysql.Config{
		User:                 config.Envs.DBuser,
		Passwd:               config.Envs.DBpassWord,
		Addr:                 config.Envs.DBaddress,
		DBName:               config.Envs.DBname,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// connect to the database.
	initStorage(db)

	// create a API serve instance
	server := api.NewAPIserver(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	// establishing a connection to the database.
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data base connected")
}
