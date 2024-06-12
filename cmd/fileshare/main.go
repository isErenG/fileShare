package main

import (
	"fileShare/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
