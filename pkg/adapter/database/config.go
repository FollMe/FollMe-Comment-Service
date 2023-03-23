package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// Load variables from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		host     = os.Getenv("DB_HOST")
		dbName   = os.Getenv("DB_NAME")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		port     = os.Getenv("DB_PORT")
	)
	// Generate connection string.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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
