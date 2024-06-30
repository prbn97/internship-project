package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/prbn97/internship-project/cmd/api"
	"github.com/prbn97/internship-project/configs"
	"github.com/prbn97/internship-project/db"
)

// the entry point for the API
func main() {

	// get configs
	cfg := mysql.Config{
		User:                 configs.Envs.DBuser,
		Passwd:               configs.Envs.DBpassWord,
		Addr:                 configs.Envs.DBaddress,
		DBName:               configs.Envs.DBname,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	// create the database
	db, err := db.NewMySQLstorage(cfg)
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
