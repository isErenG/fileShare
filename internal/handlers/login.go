package handlers

import (
	"fileShare/internal/auth"
	"fileShare/internal/data"
	"fileShare/internal/data/postgres/repository"
	"fmt"
	"net/http"
	"time"
)

type LoginHandler struct {
	Storage data.UserRepository
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		LoginPage(w, r)
		return
	}
	//TODO LATER: Add google OAuth
	username := r.FormValue("username")
	password := r.FormValue("password")

	_, err := h.Storage.VerifyUser(username, password)
	if err != nil {
		response := Response{}
		if err == repository.ErrUserNotFound {
			err := h.Storage.CreateUser(username, password)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			response.LoginMessage = "Incorrect password!"
			renderTemplate(w, "login.html", response)
			return
		}
	}

	token, err := auth.CreateToken(username)
	if err != nil {
		fmt.Println(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(2 * time.Hour),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
