package router

import (
	"fileShare/internal/auth"
	"net/http"
)

func JWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the token from the cookie
		cookie, err := r.Cookie("jwt")
		if err != nil {
			// No token provided
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		err = auth.VerifyToken(cookie.Value)
		if err != nil {
			// Token verification failed
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
