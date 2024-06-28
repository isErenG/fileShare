package handlers

import (
	"fileShare/internal/data"
	"fmt"
	"io"
	"net/http"
)

func DownloadFile(fileRepo data.FileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filecode := r.FormValue("file_code")
		if filecode == "" {
			http.Error(w, "Filename is required", http.StatusBadRequest)
			return
		}

		file, filename, err := fileRepo.DownloadObject(filecode)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "File not found!", http.StatusNotFound)
			return
		}

		// Set headers for file download
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

		_, err = io.Copy(w, file)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error downloading the file", http.StatusInternalServerError)
		}
	}
}
