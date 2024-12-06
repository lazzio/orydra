package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"orydra/config"
	"orydra/models"
	"orydra/pkg/dao"
	"orydra/pkg/hydra"
	"orydra/pkg/logger"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetClients returns a list of clients as a JSON response
func GetClients(w http.ResponseWriter, r *http.Request) {
	// Get clients from Hydra
	clients, err := hydra.GetAllHydraClients(context.Background())
	if err != nil {
		http.Error(w, "Error fetching clients", http.StatusInternalServerError)
		return
	}

	// Generate HTML options dynamically
	var options string = `<option value="">Select a client</option>`

	for _, client := range clients {
		options += fmt.Sprintf(`<option value="%s">%s</option>`, *client.ClientId, *client.ClientName)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(options))
}

func GetClientByID(w http.ResponseWriter, r *http.Request) {
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		http.Error(w, "Client ID manquant", http.StatusBadRequest)
		return
	}

	client, err := hydra.GetHydraClientByID(context.Background(), clientID)
	if err != nil {
		http.Error(w, "Client non trouvé", http.StatusNotFound)
		return
	}

	formHTML := fmt.Sprintf(`
		<h2 class="subtitle">Client details</h2>
		<form id="clientForm" hx-post="/api/client/%s/update" hx-trigger="submit">
			<input type="hidden" name="clientId" value="%s">
	`, *client.ClientId, *client.ClientId)

	if client.ClientId != nil {
		formHTML += createField("ClientId", "string", *client.ClientId)
	}
	if client.ClientName != nil {
		formHTML += createField("ClientName", "string", *client.ClientName)
	}
	if client.RedirectUris != nil {
		formHTML += createField("RedirectUris", "[]string", strings.Join(client.RedirectUris, ","))
	}
	if client.GrantTypes != nil {
		formHTML += createField("GrantTypes", "[]string", strings.Join(client.GrantTypes, ","))
	}
	if client.ResponseTypes != nil {
		formHTML += createField("ResponseTypes", "[]string", strings.Join(client.ResponseTypes, ","))
	}
	if client.Scope != nil {
		formHTML += createField("Scope", "string", *client.Scope)
	}
	if client.PolicyUri != nil {
		formHTML += createField("PolicyUri", "string", *client.PolicyUri)
	}
	if client.SkipConsent != nil {
		formHTML += createField("SkipConsent", "bool", fmt.Sprintf("%t", *client.SkipConsent))
	}
	if client.SkipLogoutConsent != nil {
		formHTML += createField("SkipLogoutConsent", "bool", fmt.Sprintf("%t", *client.SkipLogoutConsent))
	}
	if client.AccessTokenStrategy != nil {
		formHTML += createField("AccessTokenStrategy", "string", *client.AccessTokenStrategy)
	}
	if client.AllowedCorsOrigins != nil {
		formHTML += createField("AllowedCorsOrigins", "[]string", strings.Join(client.AllowedCorsOrigins, ","))
	}
	if client.Audience != nil {
		formHTML += createField("Audience", "[]string", strings.Join(client.Audience, ","))
	}
	if client.AuthorizationCodeGrantAccessTokenLifespan != nil {
		formHTML += createField("AuthorizationCodeGrantAccessTokenLifespan", "string", *client.AuthorizationCodeGrantAccessTokenLifespan)
	}
	if client.AuthorizationCodeGrantIdTokenLifespan != nil {
		formHTML += createField("AuthorizationCodeGrantIdTokenLifespan", "string", *client.AuthorizationCodeGrantIdTokenLifespan)
	}
	if client.AuthorizationCodeGrantRefreshTokenLifespan != nil {
		formHTML += createField("AuthorizationCodeGrantRefreshTokenLifespan", "string", *client.AuthorizationCodeGrantRefreshTokenLifespan)
	}
	if client.BackchannelLogoutSessionRequired != nil {
		formHTML += createField("BackchannelLogoutSessionRequired", "bool", fmt.Sprintf("%t", *client.BackchannelLogoutSessionRequired))
	}
	if client.BackchannelLogoutUri != nil {
		formHTML += createField("BackchannelLogoutUri", "string", *client.BackchannelLogoutUri)
	}
	if client.ClientCredentialsGrantAccessTokenLifespan != nil {
		formHTML += createField("ClientCredentialsGrantAccessTokenLifespan", "string", *client.ClientCredentialsGrantAccessTokenLifespan)
	}
	if client.ClientSecretExpiresAt != nil {
		formHTML += createField("ClientSecretExpiresAt", "int64", fmt.Sprintf("%d", *client.ClientSecretExpiresAt))
	}
	if client.ClientUri != nil {
		formHTML += createField("ClientUri", "string", *client.ClientUri)
	}
	if client.Contacts != nil {
		formHTML += createField("Contacts", "[]string", strings.Join(client.Contacts, ","))
	}
	if client.CreatedAt != nil {
		formHTML += createField("CreatedAt", "time.Time", client.CreatedAt.Format(time.RFC3339))
	}
	if client.FrontchannelLogoutSessionRequired != nil {
		formHTML += createField("FrontchannelLogoutSessionRequired", "bool", fmt.Sprintf("%t", *client.FrontchannelLogoutSessionRequired))
	}
	if client.FrontchannelLogoutUri != nil {
		formHTML += createField("FrontchannelLogoutUri", "string", *client.FrontchannelLogoutUri)
	}
	if client.ImplicitGrantAccessTokenLifespan != nil {
		formHTML += createField("ImplicitGrantAccessTokenLifespan", "string", *client.ImplicitGrantAccessTokenLifespan)
	}
	if client.ImplicitGrantIdTokenLifespan != nil {
		formHTML += createField("ImplicitGrantIdTokenLifespan", "string", *client.ImplicitGrantIdTokenLifespan)
	}
	if client.Jwks != nil {
		formHTML += createField("Jwks", "interface{}", fmt.Sprintf("%v", client.Jwks))
	}
	if client.JwksUri != nil {
		formHTML += createField("JwksUri", "string", *client.JwksUri)
	}
	if client.JwtBearerGrantAccessTokenLifespan != nil {
		formHTML += createField("JwtBearerGrantAccessTokenLifespan", "string", *client.JwtBearerGrantAccessTokenLifespan)
	}
	if client.LogoUri != nil {
		formHTML += createField("LogoUri", "string", *client.LogoUri)
	}
	if client.Metadata != nil {
		metadataJSON, _ := json.MarshalIndent(client.Metadata, "", "  ")
		formHTML += createField("Metadata", "map[string]interface{}", string(metadataJSON))
	}
	if client.Owner != nil {
		formHTML += createField("Owner", "string", *client.Owner)
	}
	if client.PostLogoutRedirectUris != nil {
		formHTML += createField("PostLogoutRedirectUris", "[]string", strings.Join(client.PostLogoutRedirectUris, ","))
	}
	if client.RefreshTokenGrantAccessTokenLifespan != nil {
		formHTML += createField("RefreshTokenGrantAccessTokenLifespan", "string", *client.RefreshTokenGrantAccessTokenLifespan)
	}
	if client.RefreshTokenGrantIdTokenLifespan != nil {
		formHTML += createField("RefreshTokenGrantIdTokenLifespan", "string", *client.RefreshTokenGrantIdTokenLifespan)
	}
	if client.RefreshTokenGrantRefreshTokenLifespan != nil {
		formHTML += createField("RefreshTokenGrantRefreshTokenLifespan", "string", *client.RefreshTokenGrantRefreshTokenLifespan)
	}
	if client.RegistrationAccessToken != nil {
		formHTML += createField("RegistrationAccessToken", "string", *client.RegistrationAccessToken)
	}
	if client.RegistrationClientUri != nil {
		formHTML += createField("RegistrationClientUri", "string", *client.RegistrationClientUri)
	}
	if client.RequestObjectSigningAlg != nil {
		formHTML += createField("RequestObjectSigningAlg", "string", *client.RequestObjectSigningAlg)
	}
	if client.RequestUris != nil {
		formHTML += createField("RequestUris", "[]string", strings.Join(client.RequestUris, ","))
	}
	if client.SectorIdentifierUri != nil {
		formHTML += createField("SectorIdentifierUri", "string", *client.SectorIdentifierUri)
	}
	if client.SubjectType != nil {
		formHTML += createField("SubjectType", "string", *client.SubjectType)
	}
	if client.TokenEndpointAuthMethod != nil {
		formHTML += createField("TokenEndpointAuthMethod", "string", *client.TokenEndpointAuthMethod)
	}
	if client.TokenEndpointAuthSigningAlg != nil {
		formHTML += createField("TokenEndpointAuthSigningAlg", "string", *client.TokenEndpointAuthSigningAlg)
	}
	if client.TosUri != nil {
		formHTML += createField("TosUri", "string", *client.TosUri)
	}
	if client.UpdatedAt != nil {
		formHTML += createField("UpdatedAt", "time.Time", client.UpdatedAt.Format(time.RFC3339))
	}
	if client.AccessTokenStrategy != nil {
		formHTML += createField("AccessTokenStrategy", "string", *client.AccessTokenStrategy)
	}

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

func createField(name, typ, value string) string {
	if typ == "bool" {
		checked := ""

		if value == "true" {
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
	}

	return fmt.Sprintf(`
		<div class="field">
			<label><strong>%s</strong> (%s)</label>
			<div class="control">
				<input id="%s" name="%s" class="input" type="text" value="%s">
			</div>
		</div>
	`, name, typ, name, name, value)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	envVars := config.SetEnv()

	// Get client ID from URL
	clientID := chi.URLParam(r, "id")
	if clientID == "" {
		logger.Logger.Error("Client ID missing", "error", "No client ID provided")
		http.Error(w, "Client ID missing", http.StatusBadRequest)
		return
	}

	// Get existing client
	var client models.Client
	if err := dao.PgDb.Table(envVars.POSTGRES_CLIENT_TABLE).Where("id = ?", clientID).First(&client).Error; err != nil {
		logger.Logger.Error("Client not found", "error", err)
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	// Parse the form
	if err := r.ParseForm(); err != nil {
		logger.Logger.Error("Error parsing the form", "error", err)
		http.Error(w, "Error processing the form", http.StatusBadRequest)
		return
	}

	// Update the client fields with the values from the form
	clientType := reflect.TypeOf(client)
	clientValue := reflect.ValueOf(&client).Elem()

	for i := 0; i < clientType.NumField(); i++ {
		field := clientType.Field(i)
		formValue := r.FormValue(field.Name)

		if formValue != "" {
			fieldValue := clientValue.Field(i)

			switch field.Type.Kind() {
			case reflect.String:
				fieldValue.SetString(formValue)
			case reflect.Bool:
				fieldValue.SetBool(formValue == "on" || formValue == "true")
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					values := strings.Split(formValue, ",")
					fieldValue.Set(reflect.ValueOf(values))
				}
			case reflect.Int32:
				intValue, err := strconv.ParseInt(formValue, 10, 32)
				if err != nil {
					logger.Logger.Error("Error parsing int32", "error", err)
				}
				fieldValue.SetInt(intValue)
			}

			// Manage time.Time fields
			if field.Type == reflect.TypeOf(time.Time{}) {
				timeValue, err := time.Parse(time.RFC3339, formValue)
				if err != nil {
					logger.Logger.Error("Error parsing time", "error", err)
				}
				fieldValue.Set(reflect.ValueOf(timeValue))
			}

			// Manage uuid.UUID fields
			if field.Type == reflect.TypeOf(uuid.UUID{}) {
				uuidValue, err := uuid.Parse(formValue)
				if err != nil {
					logger.Logger.Error("Error parsing UUID", "error", err)
				}
				fieldValue.Set(reflect.ValueOf(uuidValue))
			}

			// Manage sql.NullInt64 fields
			if field.Type == reflect.TypeOf(sql.NullInt64{}) {
				nullIntValue, err := strconv.ParseInt(formValue, 10, 64)
				if err != nil {
					logger.Logger.Error("Error parsing SQL Null Int64", "error", err)
					continue
				}
				// Create a new sql.NullInt64 structure
				newNullInt := sql.NullInt64{
					Int64: nullIntValue,
					Valid: true,
				}
				fieldValue.Set(reflect.ValueOf(newNullInt))
			}

			// Manage []byte fields
			if field.Type == reflect.TypeOf([]byte{}) {
				// Convert the string to a slice
				values := strings.Split(formValue, ",")
				// Serialize to JSON
				jsonData, err := json.Marshal(values)
				if err != nil {
					logger.Logger.Error("Error marshaling JSON", "error", err)
					continue
				}
				fieldValue.SetBytes(jsonData)
			}
		}
	}

	// Save the changes to the database
	if err := dao.PgDb.Table(envVars.POSTGRES_CLIENT_TABLE).Save(&client).Error; err != nil {
		logger.Logger.Error("Error updating the client", "error", err)
		http.Error(w, "Error updating the client", http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(`<div class="notification is-danger">Error updating the client %s with ID %s: %s</div>`, client.ClientName, clientID, err)))
		return
	}

	// Redirect to the index page and display a success message
	http.Redirect(w, r, "/", http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<div class="notification is-success">Client %s with ID %s updated successfully</div>`, client.ClientName, clientID)))
}
