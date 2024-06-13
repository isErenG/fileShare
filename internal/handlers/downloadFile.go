package handlers

import (
	"fileShare/internal/di"
	"io"
	"net/http"
	"os"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	filecode := r.URL.Query().Get("filename")
	if filecode == "" {
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	repo := di.GetRepository()
	file, err := repo.RetrieveFile(filecode)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error downloading the file", http.StatusInternalServerError)
	}
}
