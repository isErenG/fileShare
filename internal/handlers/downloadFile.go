package handlers

import (
	"fileShare/internal/data"
	"io"
	"net/http"
)

func DownloadFile(fileRepo data.FileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filecode := r.FormValue("file_code")

		response := Response{}

		if filecode == "" {
			response.DownloadMessage = "File code is required!"
			renderTemplate(w, "index.html", response)
			return
		}

		file, filename, contentType, err := fileRepo.DownloadFile(filecode)
		if err != nil {
			response.DownloadMessage = "File code is not found!"
			renderTemplate(w, "index.html", response)
			return
		}

		// Set headers for file download
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		w.Header().Set("Content-Type", contentType)

		_, err = io.Copy(w, file)
		if err != nil {
			response.DownloadMessage = "Internal server error! Contact administrator."
		}

		response.DownloadMessage = "Successfully downloaded!"
		renderTemplate(w, "index.html", response)

	}
}
