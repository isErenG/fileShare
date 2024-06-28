package s3

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type MinIOClient struct {
	client *minio.Client
	bucket string
}

func NewMinIOClient() (*MinIOClient, error) {
	if err := godotenv.Load("/app/.env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	endpoint := os.Getenv("ENDPOINT")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("USE_SSL"))
	bucket := os.Getenv("BUCKET")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	exists, err := minioClient.BucketExists(context.Background(), bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check if bucket exists: %w", err)
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}
	}

	return &MinIOClient{
		client: minioClient,
		bucket: bucket,
	}, nil
}

func (c *MinIOClient) UploadObject(key string, file multipart.File, fileSize int64, contentType string, originalFilename string) error {
	_, err := c.client.PutObject(context.Background(), c.bucket, key, file, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("File has been uploaded to %s. Key: %s, Filesize: %s, Filename: %s\n", c.bucket, key, fileSize, originalFilename)
	return nil
}

func (c *MinIOClient) DownloadObject(key string) (io.Reader, error) {
	object, err := c.client.GetObject(context.Background(), c.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("Error fetching object '%s' from bucket '%s': %s\n", key, c.bucket, err.Error())
		return nil, err
	}

	return object, nil
}
