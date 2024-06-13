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
	r.HandleFunc("POST /login", handlers.Login)

	http.Handle("/", r)

	// Serve static files
	staticFs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/", http.StripPrefix("/static", staticFs))

	return r
}
