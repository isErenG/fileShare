package handlers

import (
	"fileShare/internal/data"
	"fileShare/internal/di"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

}
