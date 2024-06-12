package data

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalFileRepository struct {
	basePath string
}

func NewLocalFileRepository(basePath string) *LocalFileRepository {
	return &LocalFileRepository{basePath: basePath}
}

func (repo *LocalFileRepository) RetrieveFile(filename string) (io.ReadCloser, error) {
	path := filepath.Join(repo.basePath, filename)
	fmt.Println("Retrieving file:", path)
	return os.Open(path)
}

func (repo *LocalFileRepository) SaveFile(filename string, file io.Reader) error {
	path := filepath.Join(repo.basePath, filename)
	fmt.Println("Saving file:", path)
	dest, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, file)
	return err
}
