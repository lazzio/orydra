package handlers

import (
	"fmt"
	"net/http"
	"orydra/pkg/logger"

	"orydra/pkg/hydra"
)

func UpdateClientForm(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	if err := r.ParseForm(); err != nil {
		msg := fmt.Sprintf("Error parsing the form: %s", err)
		logger.Logger.Error("Error parsing the form", "error", err)

		w.Write([]byte(fmt.Sprintf(`<div class="notification is-error">
		Error while parsing the form. Please try again.
		<br>
		Cause : %s
		</div>`, msg)))
		return
	}

	clientID := r.Form.Get("clientID")
	if clientID == "" {
		msg := "No client ID provided"
		logger.Logger.Error("Client ID missing", "error", msg)

		w.Write([]byte(fmt.Sprintf(`<div class="notification is-error">
		Error : %s</div>`, msg)))
		return
	}

	// Update the client
	err := hydra.UpdateOAuth2ClientUsingJsonPatch(clientID, r.Form)
	if err != nil {
		msg := fmt.Sprintf("Error updating the client: %s", err)
		logger.Logger.Error("Error updating the client", "error", err)

		w.Write([]byte(fmt.Sprintf(`<div class="notification is-error">
		Error updating client with ID %s. Please try again.
		<br>
		Cause : %s
		</div>`, clientID, msg)))
		return
	}

	// Redirect to the index page and display a success message
	http.Redirect(w, r, "/", http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<div class="notification is-success">Client %s with ID %s updated successfully</div>`, r.Form.Get("ClientName"), clientID)))
}
