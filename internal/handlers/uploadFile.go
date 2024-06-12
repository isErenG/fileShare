package handlers

import (
	"encoding/json"
	"fileShare/internal/di"
	"fileShare/pkg/filecodes"
	"fmt"
	"net/http"
)

type response struct {
	FileCode string `json:"file-code"`
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	code := filecodes.AddFileCode(handler.Filename)
	fmt.Println(code)
	// Add err checking

	repo := di.GetRepository()
	err = repo.SaveFile(handler.Filename, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Construct the response object
	response := response{
		FileCode: code,
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)

	//fmt.Fprintf(w, "File uploaded successfully: %v", handler.Filename)
}
