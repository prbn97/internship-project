package main

// we need to create tables in our database, like users and tasks, this is a way.

// holding the history of creation and changes of the database in the folder ./migrations.

import (
	"log"
	"os"

	"github.com/prbn97/internship-project/configs"
	"github.com/prbn97/internship-project/db"

	mysqlDriver "github.com/go-sql-driver/mysql"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg := mysqlDriver.Config{
		User:                 configs.Envs.DBuser,
		Passwd:               configs.Envs.DBpassWord,
		Addr:                 configs.Envs.DBaddress,
		DBName:               configs.Envs.DBname,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLstorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// returns a new Migrate instance from a source URL
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Version returns the currently active migration version.
	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	// Up looks at the currently active migration version and will migrate all the way up (applying all up migrations).
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	// Down looks at the currently active migration version and will migrate all the way down (applying all down migrations).
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

}
