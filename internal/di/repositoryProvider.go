package di

import "fileShare/internal/data"

func GetRepository() *data.LocalFileRepository {
	return data.NewLocalFileRepository("/Users/ereng/GolandProjects/fileShare/test/")
}
