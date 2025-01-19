package handlers

import (
	"net/http"

	"orydra/core/templates"
)

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Manage Ory Hydra Client details",
	}

	templates.RenderTemplate(w, "update-client", data)
}
