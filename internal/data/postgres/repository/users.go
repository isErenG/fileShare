package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct PostgreSQL connection string
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dbname, password, host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		fmt.Println(12)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(15)
		fmt.Println(err)
		return nil, err
	}

	postgresStore := new(PostgresStore)
	postgresStore.db = db

	if err = postgresStore.init(); err != nil {
		fmt.Println(18)
		fmt.Println(err)
		return nil, err
	}

	return postgresStore, nil
}

func (s *PostgresStore) init() error {
	err := s.createAccountTable()
	if err != nil {
		fmt.Println(20)
		fmt.Println(err)
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
		fmt.Println(1)
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *PostgresStore) CreateUser(username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err = s.db.Exec(query, username, hashedPassword)
	if err != nil {
		fmt.Println(5423)
		fmt.Println(err)
		return err
	}

	fmt.Println("user added!")
	return nil
}

func (s *PostgresStore) GetUserByUsername(username string) (*User, error) {
	query := `SELECT username, password FROM users WHERE username=$1`

	rows, err := s.db.Query(query, username)
	if err != nil {
		fmt.Println(1434)
		fmt.Println(err)
		return nil, err
	}

	user := new(User)

	rows.Next()
	err = rows.Scan(&user.Username, &user.Password)
	if err != nil {
		fmt.Println(1666)
		fmt.Println(err)
		return nil, ErrUserNotFound
	}

	fmt.Println("user retrieved")
	return user, nil
}

func (s *PostgresStore) DeleteUserByID(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		fmt.Println(1e34)
		fmt.Println(err)
		return err
	}

	fmt.Println("user deleted!")
	return nil
}

func (s *PostgresStore) VerifyUser(username, password string) (bool, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Password does not match
		return false, errors.New("incorrect password")
	}

	// Password matches
	return true, nil
}
