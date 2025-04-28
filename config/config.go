package config

import "os"

type Config struct {
	MongoURI    string
	MongoDB     string
	SMTPHost    string
	SMTPPort    string
	SMTPUser    string
	SMTPPass    string
	SMTPFrom    string
	NATSUrl     string
	ServicePort string
}

func New() Config {
	return Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDB:     os.Getenv("MONGO_DB"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPUser:    os.Getenv("SMTP_USER"),
		SMTPPass:    os.Getenv("SMTP_PASS"),
		SMTPFrom:    os.Getenv("SMTP_FROM"),
		NATSUrl:     os.Getenv("NATS_URL"),
		ServicePort: os.Getenv("SERVICE_PORT"),
	}
}
