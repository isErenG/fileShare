package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostgresStore struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewPostgresStore() (*PostgresStore, error) {
	// Load environment variables from .env file
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct PostgreSQL connection string
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	err := s.createAccountTable()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query :=
		`CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	username VARCHAR(20) NOT NULL,
    	password TEXT NOT NULL);`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) CreateUser(username string, password string) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err := s.db.Exec(query, username, password)
	if err != nil {
		return err
	}

	fmt.Println("user added!")
	return nil
}

func (s *PostgresStore) GetUserByUsername(username string) (*User, error) {
	query := `SELECT username FROM users WHERE username=$1`

	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}

	user := new(User)

	rows.Next()
	err = rows.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, nil
	}

	fmt.Println("user retrieved")
	return user, nil
}

func (s *PostgresStore) DeleteUserByID(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	fmt.Println("user deleted!")
	return nil
}
