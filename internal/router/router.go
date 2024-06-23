package router

import (
	"fileShare/internal/handlers"
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Apply JWT authorization middleware to specific routes
	authHome := JWTAuthorization(http.HandlerFunc(handlers.Home))
	r.Handle("/", authHome)

	r.HandleFunc("/upload", handlers.UploadFile)
	r.HandleFunc("/download", handlers.DownloadFile)
	r.HandleFunc("/login", handlers.Login)

	// Serve static files
	staticFs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/", http.StripPrefix("/static", staticFs))

	return r
}
