package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string

	WSToken string
}

var AppConfig = Config{}

func init() {
	err := godotenv.Load(".env")
	if err == nil {
		log.Print("Loading variable from .env file error")
	}
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")

	AppConfig.WSToken = os.Getenv("WS_TOKEN")
}
