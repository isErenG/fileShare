package data

import "io"

type FileRepository interface {
	RetrieveFile(filename string) (io.ReadCloser, error)
	SaveFile(filename string, file io.Reader) error
}

type UserRepository interface {
	CreateUser(string, string) error
	GetUserByUsername(string) (*User, error)
}
