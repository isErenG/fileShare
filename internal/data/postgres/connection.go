package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"os"
)

// Connection struct holds the database connection
type Connection struct {
	DB *sql.DB
}

// GetNewConnection initializes a new database connection
func GetNewConnection() (*Connection, error) {
	// Load environment variables from .env file
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	// Construct PostgreSQL connection string
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)
	log.Printf("Connecting to database with connection string: %s", connStr)

	var db *sql.DB
	var err error

	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	connection := &Connection{DB: db}
	return connection, nil
}
