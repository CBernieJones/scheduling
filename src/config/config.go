package config

import (
	"log"
	"os"
)

func Load() string {
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	return dbURL
}

func Port() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatal("SERVER_PORT not set")
	}

	return port
}
