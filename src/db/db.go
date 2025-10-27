package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect(dbURL string) {
	var err error
	DB, err = sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Can't connect to db: %v", err)
	}

	log.Println("Connected to db correctly")
}
