package handlers

import (
	"fileShare/internal/di"
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	repo := di.GetRepository()
	err = repo.SaveFile(handler.Filename, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %v", handler.Filename)
}
