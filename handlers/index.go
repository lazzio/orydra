package handlers

import (
	"net/http"
	"text/template"

	"orydra/core/templates"
)

var (
	Templates *template.Template
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Home",
	}

	templates.RenderTemplate(w, "index", data)
}
