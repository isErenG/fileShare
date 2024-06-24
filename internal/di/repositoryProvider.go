package di

import (
	"fileShare/internal/data"
	mongo "fileShare/internal/data/mongo/repository"
	postgres "fileShare/internal/data/postgres/repository"
)

func GetFileRepository() (data.FileRepository, error) {
	return mongo.NewMongoFileRepository("/Users/ereng/GolandProjects/fileShare/files/")
}

func GetUserRepository() (data.UserRepository, error) {
	return postgres.NewPostgresStore()
}
