package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadFromEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println("envs are not been loaded from .env file")
	}
}
