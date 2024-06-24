package handlers

import (
	"fileShare/internal/auth"
	"fileShare/internal/data"
	"fileShare/internal/data/postgres/repository"
	"fmt"
	"net/http"
	"time"
)

type UserRepository struct {
	Storage data.UserRepository
}

func (h *UserRepository) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		LoginPage(w, r)
		return
	}
	//TODO: Add postgres db and check if password hashes match for logging
	//TODO LATER: Add google OAuth
	username := r.FormValue("username")
	password := r.FormValue("password")

	usr, err := h.Storage.GetUserByUsername(username)
	if err != nil {
		if err != repository.ErrUserNotFound {
			http.Error(w, "Password is incorrect!", http.StatusUnauthorized)
			return
		}

		err := h.Storage.CreateUser(username, password)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(usr)

	// check password hash

	token, err := auth.CreateToken(username)
	if err != nil {
		fmt.Println("Yo")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
