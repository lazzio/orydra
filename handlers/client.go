package handlers

import (
	"context"
	"fmt"
	"net/http"
	"orydra/pkg/hydra"
	"orydra/pkg/logger"
	"sort"
	"strings"

	"github.com/go-chi/chi/v5"
)

// GetClients handles HTTP requests to retrieve OAuth2 clients from Hydra.
// It fetches all clients from Hydra, sorts them by client name alphabetically,
// and returns an HTML string containing <option> elements for use in a select dropdown.
//
// The function:
// 1. Retrieves clients from Hydra using hydra.GetAllHydraClients
// 2. Sorts clients alphabetically by their ClientName
// 3. Generates HTML options with client ID as value and client name as display text
// 4. Returns the HTML string with Content-Type set to "text/html"
//
// If an error occurs while fetching clients, it logs the error and returns a 500 status code.
func GetClients(w http.ResponseWriter, r *http.Request) {
	// Get clients from Hydra
	clients, err := hydra.GetAllHydraClients(r.Context())
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la récupération des clients", "error", err, "function", funcName)
		http.Error(w, "Error fetching clients", http.StatusInternalServerError)
		return
	}

	// Sort clients by ClientName
	sort.Slice(clients, func(i, j int) bool {
		return *clients[i].ClientName < *clients[j].ClientName
	})

	// Generate HTML options dynamically
	var options string = `<option value="">Select a client</option>`

	for _, client := range clients {
		options += fmt.Sprintf(`<option value="%s">%s</option>`, *client.ClientId, *client.ClientName)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(options))
}

// GetClientByID returns a client by ID
// @return models.Client, error
func GetClientByID(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		logger.Logger.Error("Client ID manquant")
		http.Error(w, "Client ID manquant", http.StatusBadRequest)
		return
	}

	client, err := hydra.GetHydraClientByID(context.Background(), clientID)
	if err != nil {
		logger.Logger.Error("Client non trouvé", "error", err)
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	formHTML := fmt.Sprintf(`
		<h2 class="subtitle">Client details</h2>
		<form id="clientForm" hx-post="/api/client/%s/update" hx-trigger="submit">
	`, *client.ClientId)

	// Send initial client data to the form to compare with updated values
	//formHTML += fmt.Sprintf(`<input type="hidden" name="initialClient" value="%v">`, client)
	formHTML += fmt.Sprintf(`<input type="hidden" name="clientID" value="%s">`, clientID)

	// Generate form fields for each client field
	formHTML += createField("AccessTokenStrategy", "string", client.GetAccessTokenStrategy())
	formHTML += createField("AllowedCorsOrigins", "[]string", client.GetAllowedCorsOrigins())
	formHTML += createField("Audience", "[]string", client.GetAudience())
	formHTML += createField("AuthorizationCodeGrantAccessTokenLifespan", "string", client.GetAuthorizationCodeGrantAccessTokenLifespan())
	formHTML += createField("AuthorizationCodeGrantIdTokenLifespan", "string", client.GetAuthorizationCodeGrantIdTokenLifespan())
	formHTML += createField("AuthorizationCodeGrantRefreshTokenLifespan", "string", client.GetAuthorizationCodeGrantRefreshTokenLifespan())
	formHTML += createField("BackchannelLogoutSessionRequired", "bool", client.GetBackchannelLogoutSessionRequired())
	formHTML += createField("BackchannelLogoutUri", "string", client.GetBackchannelLogoutUri())
	formHTML += createField("ClientCredentialsGrantAccessTokenLifespan", "string", client.GetClientCredentialsGrantAccessTokenLifespan())
	formHTML += createField("ClientId", "string", client.GetClientId())
	formHTML += createField("ClientName", "string", client.GetClientName())
	// formHTML += createField("ClientSecretExpiresAt", "string", client.GetClientSecretExpiresAt())
	formHTML += createField("ClientUri", "string", client.GetClientUri())
	formHTML += createField("Contacts", "[]string", client.GetContacts())
	formHTML += createField("CreatedAt", "time", client.GetCreatedAt())
	formHTML += createField("FrontchannelLogoutSessionRequired", "bool", client.GetFrontchannelLogoutSessionRequired())
	formHTML += createField("FrontchannelLogoutUri", "string", client.GetFrontchannelLogoutUri())
	formHTML += createField("GrantTypes", "[]string", client.GetGrantTypes())
	formHTML += createField("ImplicitGrantAccessTokenLifespan", "string", client.GetImplicitGrantAccessTokenLifespan())
	formHTML += createField("ImplicitGrantIdTokenLifespan", "string", client.GetImplicitGrantIdTokenLifespan())
	formHTML += createField("Jwks", "interface{}", client.GetJwks())
	formHTML += createField("JwksUri", "string", client.GetJwksUri())
	formHTML += createField("JwtBearerGrantAccessTokenLifespan", "string", client.GetJwtBearerGrantAccessTokenLifespan())
	formHTML += createField("LogoUri", "string", client.GetLogoUri())
	formHTML += createField("Metadata", "interface{}", client.GetMetadata())
	formHTML += createField("Owner", "string", client.GetOwner())
	formHTML += createField("PolicyUri", "string", client.GetPolicyUri())
	formHTML += createField("PostLogoutRedirectUris", "[]string", client.GetPostLogoutRedirectUris())
	formHTML += createField("RedirectUris", "[]string", client.GetRedirectUris())
	formHTML += createField("RefreshTokenGrantAccessTokenLifespan", "string", client.GetRefreshTokenGrantAccessTokenLifespan())
	formHTML += createField("RefreshTokenGrantIdTokenLifespan", "string", client.GetRefreshTokenGrantIdTokenLifespan())
	formHTML += createField("RefreshTokenGrantRefreshTokenLifespan", "string", client.GetRefreshTokenGrantRefreshTokenLifespan())
	formHTML += createField("RegistrationAccessToken", "string", client.GetRegistrationAccessToken())
	formHTML += createField("RegistrationClientUri", "string", client.GetRegistrationClientUri())
	formHTML += createField("RequestObjectSigningAlg", "string", client.GetRequestObjectSigningAlg())
	formHTML += createField("RequestUris", "string", client.GetRequestUris())
	formHTML += createField("ResponseTypes", "string", client.GetResponseTypes())
	formHTML += createField("Scope", "string", client.GetScope())
	formHTML += createField("SectorIdentifierUri", "string", client.GetSectorIdentifierUri())
	formHTML += createField("SkipConsent", "bool", client.GetSkipConsent())
	formHTML += createField("SkipLogoutConsent", "bool", client.GetSkipLogoutConsent())
	formHTML += createField("SubjectType", "string", client.GetSubjectType())
	formHTML += createField("TokenEndpointAuthMethod", "string", client.GetTokenEndpointAuthMethod())
	formHTML += createField("TokenEndpointAuthSigningAlg", "string", client.GetTokenEndpointAuthSigningAlg())
	formHTML += createField("TosUri", "string", client.GetTosUri())
	formHTML += createField("UpdatedAt", "time", client.GetUpdatedAt())
	formHTML += createField("UserinfoSignedResponseAlg", "string", client.GetUserinfoSignedResponseAlg())

	formHTML += `<div class="field is-grouped">
		<p class="control">
			<button class="button is-primary is-rounded" type="submit">Update</button>
		</p>
		<p class="control">
			<a class="button is-danger is-rounded" href="/">Cancel</a>
		</p>
	</div></form>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formHTML))
}

// createField generates an HTML form field element based on the provided parameters.
// It creates either a checkbox field for boolean types or a text input field for other types.
//
// Parameters:
//   - name: the name/id attribute of the form field
//   - typ: the data type of the field ("bool" for checkbox, any other value for text input)
//   - value: the current value of the field (used for checked state in checkbox or value in text input)
//
// Returns a string containing the HTML markup for the form field with Bulma CSS classes.
func createField(name, typ, value any) string {
	switch typ {
	case "bool":
		checked := ""

		if value == true {
			checked = "checked"
		}

		return fmt.Sprintf(`
			<div class="field">
				<label class="checkbox"><strong>%s</strong> (%s)</label>
				<div class="control">
					<input id="%s" name="%s" type="checkbox" %s>
				</div>
			</div>
		`, name, typ, name, name, checked)

	case "[]string":
		// Convert []string to comma-separated string
		valueStr := strings.Join(value.([]string), ",")
		// Remove empty spaces from the string
		valueStr = strings.ReplaceAll(valueStr, " ", "")

		return fmt.Sprintf(`
			<div class="field">
				<label><strong>%s</strong> (%s converted to comma-separated values)</label>
				<div class="control">
					<input id="%s" name="%s" class="input" type="text" value="%s">
				</div>
			</div>
		`, name, typ, name, name, valueStr)

	default:
		return fmt.Sprintf(`
				<div class="field">
					<label><strong>%s</strong> (%s)</label>
					<div class="control">
						<input id="%s" name="%s" class="input" type="text" value="%s">
					</div>
				</div>
			`, name, typ, name, name, value)
	}
}
