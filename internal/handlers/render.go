package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("/app/templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	fmt.Println(tmplPath)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error parsing templates", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error executing templates", http.StatusInternalServerError)
	}
}
