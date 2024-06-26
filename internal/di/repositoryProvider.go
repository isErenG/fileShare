package di

import (
	"fileShare/internal/data"
	postgres "fileShare/internal/data/postgres/repository"
	"fileShare/internal/s3"
)

func GetFileRepository() (data.FileRepository, error) {
	return s3.NewMinIOClient()
}

func GetUserRepository() (data.UserRepository, error) {
	return postgres.NewPostgresStore()
}
