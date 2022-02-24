package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvVariables load environment variables
func EnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
