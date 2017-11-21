package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shumkovdenis/example"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "./database.sqlite")
	if err != nil {
		log.Panic(err)
	}

	userStore := example.NewUserStore(db)

	server := example.NewServer(userStore)

	if err := example.NewAPI(server); err != nil {
		log.Panic(err)
	}
}
