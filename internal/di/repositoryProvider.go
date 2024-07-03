package di

import (
	"fileShare/internal/data"
	db "fileShare/internal/data/postgres"
	"fileShare/internal/data/postgres/repository"
)

func GetFileRepository(connection *db.Connection) (data.FileRepository, error) {
	return repository.NewFilesStorage(connection)
}

func GetUserRepository(connection *db.Connection) (data.UserRepository, error) {
	return repository.NewUsersStorage(connection)
}
