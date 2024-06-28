package router

import (
	"fileShare/internal/data/postgres"
	"fileShare/internal/di"
	"fileShare/internal/handlers"
	"fileShare/pkg/auth/jwt"
	"fmt"
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	conn, err := db.GetNewConnection()
	if err != nil {
		fmt.Println("Error creating connection! " + err.Error())
	}

	userRepo := di.GetUserRepository(conn)
	fileRepo := di.GetFileRepository(conn)

	// Golang constructor
	lh := &handlers.LoginHandler{Storage: userRepo}

	// Apply JWT authorization middleware to specific routes
	authHome := jwt.JWTAuthorization(http.HandlerFunc(handlers.Home))
	authUpload := jwt.JWTAuthorization(handlers.UploadFile(fileRepo))
	authDownload := jwt.JWTAuthorization(handlers.DownloadFile(fileRepo))

	r.Handle("/", authHome)
	r.Handle("/upload", authUpload)
	r.Handle("/download", authDownload)

	r.HandleFunc("/login", lh.Login)

	// Serve static files
	staticFs := http.FileServer(http.Dir("/app/static"))
	r.Handle("/static/", http.StripPrefix("/static", staticFs))

	return r
}
