package handlers

import (
	"net/http"
	"text/template"

	"orydra/core/templates"
)

var (
	Templates *template.Template
)

// func init() {
// 	// Templates loading
// 	Templates = template.Must(template.ParseFiles("templates/index.html"))
// }

func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Manage Ory Hydra Client details",
	}

	templates.RenderTemplate(w, "index", data)
}
