package di

import "fileShare/internal/data"

func GetFileRepository() data.FileRepository {
	return data.NewLocalFileRepository("/Users/ereng/GolandProjects/fileShare/files/")
}

func GetUserRepository() data.UserRepository {
	return data.NewLocalUserRepository("d")
}
