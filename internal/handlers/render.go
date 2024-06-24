package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test1")
	renderTemplate(w, "index.html", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test2")
	renderTemplate(w, "login.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("/app/templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	fmt.Println(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing templates", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing templates", http.StatusInternalServerError)
	}
}
