package handlers

import (
	"net/http"
	"text/template"

	"orydra/core/templates"
)

var (
	Templates *template.Template
)

func init() {
	// Templates loading
	Templates = template.Must(template.ParseFiles("templates/index.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "index", nil)
}
