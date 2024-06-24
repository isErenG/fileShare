package router

import (
	"fileShare/internal/di"
	"fileShare/internal/handlers"
	"fileShare/pkg/auth/jwt"
	"net/http"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	userRepo, err := di.GetUserRepository()
	if err != nil {
		panic(err)
	}

	fileRepo, err := di.GetFileRepository()
	if err != nil {
		panic(err)
	}

	// Passing user repository to the login handler or smtg idk it works
	h := &handlers.UserRepository{Storage: userRepo}

	// Apply JWT authorization middleware to specific routes
	authHome := jwt.JWTAuthorization(http.HandlerFunc(handlers.Home))
	authUpload := jwt.JWTAuthorization(handlers.UploadFile(fileRepo))
	authDownload := jwt.JWTAuthorization(handlers.DownloadFile(fileRepo))

	r.Handle("/", authHome)
	r.Handle("/upload", authUpload)
	r.Handle("/download", authDownload)

	r.HandleFunc("/login", h.Login)

	// Serve static files
	staticFs := http.FileServer(http.Dir("/app/static"))
	r.Handle("/static/", http.StripPrefix("/static", staticFs))

	return r
}
