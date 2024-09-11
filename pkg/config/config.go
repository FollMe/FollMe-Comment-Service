package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

func getProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for go.mod
	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); os.IsNotExist(err) {
			parent := filepath.Dir(cwd)
			if parent == cwd {
				// We've reached the root directory without finding go.mod
				return "", fmt.Errorf("go.mod not found")
			}
			cwd = parent
		} else {
			return cwd, nil
		}
	}
}

func init() {
	rootPath, err := getProjectRoot()
	if err != nil {
		log.Print("Get root path error: ", err.Error())
	}

	envPath := filepath.Join(rootPath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Print("Loading variable from .env file error")
	}
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")

	AppConfig.WSToken = os.Getenv("WS_TOKEN")
}
