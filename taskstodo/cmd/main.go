package main

import (
	"database/sql"
	"log"
	"os"

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
		Net:                  configs.Envs.Net,
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
	server := api.NewAPIserver(":"+os.Getenv("PORT"), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// establishing a connection to the database.
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Data base connected")
}
