package data

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LocalUserRepository struct {
	filePath string
}

var ErrUserNotFound = errors.New("user not found")

func (l *LocalUserRepository) CreateUser(username string, password string) error {
	// Load existing users from file
	users, err := l.loadUsersFromFile()
	if err != nil {
		return err
	}

	// Check if user already exists
	if _, ok := users[username]; ok {
		return errors.New("user already exists")
	}

	// Add the new user
	users[username] = &User{
		Username: username,
		Password: password,
	}

	// Save the updated user list to file
	err = l.saveUsersToFile(users)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocalUserRepository) GetUserByUsername(username string) (*User, error) {
	// Load existing users from file
	users, err := l.loadUsersFromFile()
	if err != nil {
		return nil, err
	}

	// Find and return the user by username
	user, ok := users[username]
	if !ok {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (l *LocalUserRepository) loadUsersFromFile() (map[string]*User, error) {
	users := make(map[string]*User)

	file, err := os.Open(l.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist yet, return empty user map
			return users, nil
		}
		return nil, err
	}
	defer file.Close()

	// Decode JSON content
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	return users, nil
}

func (l *LocalUserRepository) saveUsersToFile(users map[string]*User) error {
	file, err := os.OpenFile(l.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode users map to JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		return err
	}

	return nil
}

func NewLocalUserRepository(filePath string) *LocalUserRepository {
	return &LocalUserRepository{filePath: filePath}
}
