package di

import "fileShare/internal/data"

func GetRepository() data.FileRepository {
	return data.NewLocalFileRepository("/Users/ereng/GolandProjects/fileShare/files/")
}
