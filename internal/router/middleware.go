package router

import (
	"fileShare/internal/auth"
	"net/http"
)

func JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			// No token provided
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Extract the token value (assuming it's in the format "Bearer <token>")
		tokenString = tokenString[len("Bearer "):]

		// Verify the token
		err := auth.VerifyToken(tokenString)
		if err != nil {
			// Token verification failed
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
