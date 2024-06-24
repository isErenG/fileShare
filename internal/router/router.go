package router

import (
	"fileShare/internal/handlers"
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Apply JWT authorization middleware to specific routes
	authHome := JWTAuthorization(http.HandlerFunc(handlers.Home))
	authUpload := JWTAuthorization(http.HandlerFunc(handlers.UploadFile))
	authDownload := JWTAuthorization(http.HandlerFunc(handlers.DownloadFile))

	r.Handle("/", authHome)
	r.Handle("/upload", authUpload)
	r.Handle("/download", authDownload)

	r.HandleFunc("/login", handlers.Login)

	// Serve static files
	staticFs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/", http.StripPrefix("/static", staticFs))

	return r
}
