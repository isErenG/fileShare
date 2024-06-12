package main

import (
	"fileShare/internal/router"
	"log"
	"net/http"
)

// Development purposes only
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := router.New()

	corsR := enableCORS(r)

	if err := http.ListenAndServe(":8080", corsR); err != nil {
		log.Fatal(err)
	}
}
