package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthServiceAddr     string
	DocumentsServiceURL string
	JWTSecret           string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		AuthServiceAddr:     os.Getenv("AUTH_SERVICE_ADDR"),
		DocumentsServiceURL: os.Getenv("DOCUMENTS_SERVICE_URL"),
		JWTSecret:           os.Getenv("JWT_SECRET"),
	}
}