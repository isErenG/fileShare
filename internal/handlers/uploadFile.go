package handlers

import (
	"fileShare/internal/data"
	"fileShare/lib/filecodes"
	"fmt"
	"log"
	"net/http"
)

func UploadFile(fileRepo data.FileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		code := filecodes.CreateCode()
		// TODO: Add err checking
		fmt.Println(code)

		contentType := handler.Header.Get("Content-Type")
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		err = fileRepo.UploadFile(code, file, handler.Size, contentType, handler.Filename)
		if err != nil {
			log.Printf("Failed to upload file: %v", err)
			http.Error(w, "Failed to upload file", http.StatusInternalServerError)
			return
		}

		// Construct the response object
		response := Response{
			FileCode: code,
		}

		renderTemplate(w, "index.html", response)
	}
}
