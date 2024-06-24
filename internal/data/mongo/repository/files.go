package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"time"
)

type MongoFileRepository struct {
	client *mongo.Client
}

func NewMongoFileRepository() (*MongoFileRepository, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return &MongoFileRepository{client: client}, nil
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
