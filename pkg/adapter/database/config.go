package database

import (
	"database/sql"
	"fmt"
	"follme/comment-service/pkg/config"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// Load variables from .env
	err := godotenv.Load(".env")
	if err == nil {
		log.Print("Loading variable from .env file")
	}
	var (
		host     = config.AppConfig.DB_HOST
		dbName   = config.AppConfig.DB_NAME
		username = config.AppConfig.DB_USERNAME
		password = config.AppConfig.DB_PASSWORD
		port     = config.AppConfig.DB_PORT
	)
	// Generate connection string.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, username, password, dbName,
	)

	// Get a database handle.
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connect DB successfully!")
	return db
}
