package internal

import (
	"fileShare/internal/handlers"
	"net/http"
)

func NewRouter() error {
	http.HandleFunc("/upload", handlers.UploadFile)
	http.HandleFunc("/download", handlers.DownloadFile)

	err := http.ListenAndServe(":8080", nil)

	return err
}
