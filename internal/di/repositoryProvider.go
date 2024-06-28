package di

import (
	"fileShare/internal/data"
	db "fileShare/internal/data/postgres"
	"fileShare/internal/data/postgres/repository"
)

func GetFileRepository(connection *db.Connection) data.FileRepository {
	return repository.NewFilesStorage(connection)
}

func GetUserRepository(connection *db.Connection) data.UserRepository {
	return repository.NewUsersStorage(connection)
}
