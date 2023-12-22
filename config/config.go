package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DATABASE_URL string

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
}
