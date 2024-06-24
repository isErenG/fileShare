package data

import (
	"fileShare/internal/data/postgres/repository"
	"io"
)

type FileRepository interface {
	RetrieveFile(string) (io.ReadCloser, error)
	SaveFile(string, io.Reader) error
	DeleteFile(string) error
}

type UserRepository interface {
	CreateUser(string, string) error
	GetUserByUsername(string) (*repository.User, error)
	DeleteUserByID(int) error
	VerifyUser(string, string) (bool, error)
}
