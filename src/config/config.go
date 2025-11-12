package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() string {

	dbUser := getVariable("DB_USER")
	dbPass := getVariable("DB_PASSWORD")
	dbHost := getVariable("DB_HOST")
	dbPort := getVariable("DB_PORT")

	return fmt.Sprintf("postgres://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

}

func Port() string {
	return getVariable("SERVER_PORT")
}

func getVariable(variable string) string {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	variableText := os.Getenv(variable)

	if variableText == "" {
		log.Fatalf("%s not set", variable)
	}

	return variableText
}
