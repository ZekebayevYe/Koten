package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AuthServiceAddr         string
	NotificationServiceAddr string
	JWTSecret               string
	NewsServiceAddr         string
	ReportServiceAddr       string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		AuthServiceAddr:         os.Getenv("AUTH_SERVICE_ADDR"),
		JWTSecret:               os.Getenv("JWT_SECRET"),
		NotificationServiceAddr: os.Getenv("NOTIFICATION_SERVICE_ADDR"),
		NewsServiceAddr:         os.Getenv("NEWS_SERVICE_ADDR"),
		ReportServiceAddr:       os.Getenv("REPORT_SERVICE_ADDR"),
	}
}
