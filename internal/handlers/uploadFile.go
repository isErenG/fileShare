package handlers

import (
	"fileShare/internal/di"
	"fileShare/pkg/filecodes"
	"fmt"
	"net/http"
)

type response struct {
	FileCode string `json:"file_code"`
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	code := filecodes.AddFileCode(handler.Filename)
	// TODO: Add err checking
	fmt.Println(code)

	repo := di.GetFileRepository()
	err = repo.SaveFile(handler.Filename, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Construct the response object
	response := response{
		FileCode: code,
	}

	if err != nil {
		panic(err)
	}

	renderTemplate(w, "index.html", response)

}
