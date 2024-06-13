package router

import (
	"fileShare/internal/handlers"
	"net/http"
)

func New() *http.ServeMux {

	r := http.NewServeMux()

	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("POST /upload", handlers.UploadFile)
	r.HandleFunc("GET /download", handlers.DownloadFile)

	http.Handle("/", r)

	return r
}
