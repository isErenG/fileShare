package handlers

import (
	"fileShare/internal/auth"
	"fileShare/internal/data"
	"fileShare/internal/di"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		LoginPage(w, r)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	repo := di.GetUserRepository()

	usr, err := repo.GetUserByUsername(username)
	if err != nil {
		if err != data.ErrUserNotFound {
			// TODO: Return password is wrong!
			http.Error(w, "Password is incorrect!", http.StatusUnauthorized)
			return
		}

		repo.CreateUser(username, password)
	}

	fmt.Println(usr)

	// check password hash

	token, err := auth.CreateToken(username)
	if err != nil {
		fmt.Println("Yo")
	}

	// Set the token in the response header or response body as needed
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
