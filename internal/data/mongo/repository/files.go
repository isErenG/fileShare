package repository

import (
	"io"
)

type MongoFileRepository struct {
	basePath string
}

func NewMongoFileRepository(basePath string) (*MongoFileRepository, error) {
	return &MongoFileRepository{basePath: basePath}, nil
}

func (repo *MongoFileRepository) RetrieveFile(filename string) (io.ReadCloser, error) {
	return nil, nil
}

func (repo *MongoFileRepository) SaveFile(filename string, file io.Reader) error {
	return nil

}

func (repo *MongoFileRepository) DeleteFile(s string) error {
	return nil
}
