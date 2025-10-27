package config

import (
	"log"
	"os"
)

func Load() string {
	return getVariable("DATABASE_URL")
}

func Port() string {
	return getVariable("SERVER_PORT")
}

func getVariable(variable string) string {
	variableText := os.Getenv(variable)

	if variableText == "" {
		log.Fatalf("%s not set", variable)
	}

	return variableText
}
