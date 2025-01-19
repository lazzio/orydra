package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"orydra/core/templates"
	"orydra/pkg/hydra"
	"time"

	hc "github.com/ory/hydra-client-go/v2"
)

type CreateClientResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func CreateClientForm(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Create a new Ory Hydra client",
		Message: "",
	}
	templates.RenderTemplate(w, "create-client", data)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	// Gérer la requête GET pour SSE
	if r.Method == http.MethodGet {
		// Configuration SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}

		// Récupérer les données du query parameter
		clientName := r.URL.Query().Get("client_name")
		if clientName == "" {
			_, _ = json.Marshal(map[string]string{"error": "client_name is required"})
			flusher.Flush()
			return
		}

		// Créer un canal pour recevoir le résultat
		resultChan := make(chan map[string]string, 1)
		errChan := make(chan error, 1)

		newClient := &hc.OAuth2Client{
			ClientName: &clientName,
		}

		// Créer un nouveau contexte pour la goroutine
		ctx := context.Background()

		// Lancer la création du client dans une goroutine
		go func() {
			client, err := hydra.CreateNewHydraClient(ctx, newClient)
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- map[string]string{
				"clientId":     *client.ClientId,
				"clientSecret": *client.ClientSecret,
			}
		}()

		// Attendre le résultat
		select {
		case result := <-resultChan:
			_, _ = json.Marshal(result)
			flusher.Flush()
		case err := <-errChan:
			_, _ = json.Marshal(map[string]string{"error": err.Error()})
			flusher.Flush()
		case <-time.After(10 * time.Second):
			_, _ = json.Marshal(map[string]string{"error": "timeout"})
			flusher.Flush()
		}
		return
	}

	// Gérer la requête POST (pour la rétrocompatibilité si nécessaire)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Use GET method with SSE", http.StatusMethodNotAllowed)
}
