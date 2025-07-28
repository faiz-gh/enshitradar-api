package main

import (
	"database/sql"
	"log"

	"github.com/faiz-gh/enshitradar-api/cmd/api"
	"github.com/faiz-gh/enshitradar-api/config"
	"github.com/faiz-gh/enshitradar-api/db"
)

func main() {
	db, err := db.NewPostgreSQLStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":"+config.ENV.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully Connected!")
}
