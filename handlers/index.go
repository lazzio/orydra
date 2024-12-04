package handlers

import (
	"net/http"
	"text/template"
)

var (
	Templates *template.Template
)

func init() {
	// Templates loading
	Templates = template.Must(template.ParseFiles("templates/index.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := Templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Error rendering the page", http.StatusInternalServerError)
	}
}
