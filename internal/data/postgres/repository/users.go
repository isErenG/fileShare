package repository

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	db "fileShare/internal/data/postgres"
)

type UserRepository struct {
	Storage *db.Connection
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var ErrUserNotFound = errors.New("user not found")

func NewUsersStorage(conn *db.Connection) (*UserRepository, error) {
	userRepository := new(UserRepository)
	userRepository.Storage = conn

	err := userRepository.UsersInit()
	if err != nil {
		return nil, err
	}

	return userRepository, nil
}

func (s *UserRepository) UsersInit() error {
	err := s.createAccountTable()
	if err != nil {
		fmt.Println(20)
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *UserRepository) createAccountTable() error {
	query :=
		`CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	username VARCHAR(20) NOT NULL,
    	password TEXT NOT NULL);`

	_, err := s.Storage.DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *UserRepository) CreateUser(username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err = s.Storage.DB.Exec(query, username, hashedPassword)
	if err != nil {
		fmt.Println(5423)
		fmt.Println(err)
		return err
	}

	fmt.Println("user added!")
	return nil
}

func (s *UserRepository) GetUserByUsername(username string) (*User, error) {
	query := `SELECT username, password FROM users WHERE username=$1`

	row := s.Storage.DB.QueryRow(query, username)

	user := new(User)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, ErrUserNotFound
	}

	fmt.Println("user retrieved")
	return user, nil
}

func (s *UserRepository) DeleteUserByID(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := s.Storage.DB.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *UserRepository) VerifyUser(username, password string) (bool, error) {
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
