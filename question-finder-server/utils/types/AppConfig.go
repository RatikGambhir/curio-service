package util_types

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	POSTGRES_URL string
	DB_NAME      string
}

func GetAppConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	return AppConfig{
		POSTGRES_URL: os.Getenv("POSTGRES_URL"),
		DB_NAME:      os.Getenv("DB_NAME"),
	}
}
