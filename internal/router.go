package internal

import (
	"fileShare/internal/handlers"
	"net/http"
)

func newRouter() *http.ServeMux {

	r := http.NewServeMux()

	r.HandleFunc("POST /upload", handlers.UploadFile)
	r.HandleFunc("GET /download", handlers.DownloadFile)

	http.Handle("/", r)

	return r
}

func StartRouter() {
	router := newRouter()

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
