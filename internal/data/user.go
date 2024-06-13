package data

import "errors"

type User struct {
	Username string
	Password string
}

type LocalUserRepository struct {
	filePath string
}

var ErrUserNotFound = errors.New("user not found")

func (l LocalUserRepository) CreateUser(user string, pass string) error {
	return nil
}

func (l LocalUserRepository) GetUserByUsername(user string) (*User, error) {
	return nil, nil
}

func NewLocalUserRepository(filePath string) *LocalUserRepository {
	return &LocalUserRepository{filePath: filePath}
}
