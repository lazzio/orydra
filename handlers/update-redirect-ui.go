package handlers

import (
	"net/http"
	"orydra/pkg/logger"

	"github.com/go-chi/chi/v5"
)

func UpdateRedirectUI(w http.ResponseWriter, r *http.Request) {
	// Get client ID from URL
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		logger.Logger.Error("Client ID missing", "error", "No client ID provided")
		http.Error(w, "Client ID missing", http.StatusBadRequest)
		return
	}

	// Parse the form
	if err := r.ParseForm(); err != nil {
		logger.Logger.Error("Error parsing the form", "error", err)
		http.Error(w, "Error processing the form", http.StatusBadRequest)
		return
	}
}
