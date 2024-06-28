package repository

import (
	"database/sql"
	db "fileShare/internal/data/postgres"
	"fileShare/internal/s3"
	"fmt"
	"io"
	"mime/multipart"
)

type FileRepository struct {
	Storage *sql.DB
	S3      *s3.MinIOClient
}

func NewFilesStorage(conn *db.Connection) *FileRepository {
	filesRepository := new(FileRepository)
	filesRepository.Storage = conn.DB

	minIO, err := s3.NewMinIOClient()
	if err != nil {
		fmt.Println("Failed to connect to minIO: " + err.Error())
		return nil
	}

	filesRepository.S3 = minIO

	err = filesRepository.FilesInit()
	if err != nil {
		fmt.Println("Failed to initialize files repository: " + err.Error())
		return nil
	}

	return filesRepository
}

func (s *FileRepository) FilesInit() error {
	err := s.createFilesTable()
	if err != nil {
		fmt.Println("Failed to create files table: " + err.Error())
		return err
	}

	return nil
}

func (s *FileRepository) createFilesTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS files (
			slug VARCHAR(6) PRIMARY KEY,
			file_name VARCHAR(100) NOT NULL,
			content_type VARCHAR(30) NOT NULL
		);`
	_, err := s.Storage.Exec(query)
	if err != nil {
		fmt.Println("Error creating files table: " + err.Error())
		return err
	}

	return nil
}

func (s *FileRepository) UploadFile(code string, file multipart.File, size int64, contentType string, fileName string) error {
	if s.S3 == nil {
		return fmt.Errorf("S3 client is nil")
	}

	err := s.S3.UploadObject(code, file, size, contentType, fileName)
	if err != nil {
		fmt.Println("Error uploading file to S3: " + err.Error())
		return err
	}

	if s.Storage == nil {
		return fmt.Errorf("SQL storage is nil")
	}

	query := `INSERT INTO files (slug, file_name, content_type) VALUES ($1, $2, $3)`
	_, err = s.Storage.Exec(query, code, fileName, contentType)
	if err != nil {
		fmt.Println("Error inserting file metadata into database: " + err.Error())
		return err
	}

	return nil
}

func (s *FileRepository) DownloadFile(code string) (io.Reader, string, string, error) {
	if s.S3 == nil {
		return nil, "", "", fmt.Errorf("S3 client is nil")
	}

	// Query to retrieve file_name and content_type from database
	query := `SELECT file_name, content_type FROM files WHERE slug = $1`
	row := s.Storage.QueryRow(query, code)

	var fileName, contentType string
	err := row.Scan(&fileName, &contentType)
	if err != nil {
		fmt.Println("Error retrieving file metadata from database: " + err.Error())
		return nil, "", "", err
	}

	file, err := s.S3.DownloadObject(code)
	if err != nil {
		fmt.Println("Error downloading file from S3: " + err.Error())
		return nil, "", "", err
	}

	return file, fileName, contentType, nil
}
