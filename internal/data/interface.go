package data

import (
	"fileShare/internal/data/postgres/repository"
	"io"
	"mime/multipart"
)

type FileRepository interface {
	DownloadObject(string) (io.Reader, string, string, error)
	UploadObject(string, multipart.File, int64, string, string) error
}

type UserRepository interface {
	CreateUser(string, string) error
	GetUserByUsername(string) (*repository.User, error)
	DeleteUserByID(int) error
	VerifyUser(string, string) (bool, error)
}
