package handlers

import (
	"fmt"
	"net/http"
	"orydra/pkg/logger"

	"orydra/pkg/hydra"
)

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	if err := r.ParseForm(); err != nil {
		logger.Logger.Error("Error parsing the form", "error", err)
		http.Error(w, "Error processing the form", http.StatusBadRequest)
		return
	}

	// Get initial Client
	//initialClient := r.Form.Get("initialClient")
	// if initialClient == "" {
	// 	logger.Logger.Error("Initial client details missing", "error", "No initial client details provided")
	// 	http.Error(w, "Initial client details missing", http.StatusBadRequest)
	// 	return
	// }

	clientID := r.Form.Get("clientID")
	if clientID == "" {
		logger.Logger.Error("Client ID missing", "error", "No client ID provided")
		http.Error(w, "Client ID missing", http.StatusBadRequest)
		return
	}

	// Update the client fields with the values from the form

	// client := &hc.OAuth2Client{
	// 	ClientId: &clientID,
	// 	AccessTokenStrategy: hc.PtrString(r.Form.Get("AccessTokenStrategy")),
	// 	AllowedCorsOrigins:  strings.Split(r.Form.Get("AllowedCorsOrigins"), ","),
	// 	Audience: 		  strings.Split(r.Form.Get("Audience"), ","),
	// 	AuthorizationCodeGrantAccessTokenLifespan: hc.PtrString(r.Form.Get("AuthorizationCodeGrantAccessTokenLifespan")),
	// 	AuthorizationCodeGrantIdTokenLifespan: hc.PtrString(r.Form.Get("AuthorizationCodeGrantIdTokenLifespan")),
	// 	AuthorizationCodeGrantRefreshTokenLifespan: hc.PtrString(r.Form.Get("AuthorizationCodeGrantRefreshTokenLifespan")),
	// 	BackchannelLogoutSessionRequired: hc.PtrBool(r.Form.Get("BackchannelLogoutSessionRequired") == "true"),
	// 	BackchannelLogoutUri: hc.PtrString(r.Form.Get("BackchannelLogoutUri")),
	// 	ClientCredentialsGrantAccessTokenLifespan: hc.PtrString(r.Form.Get("ClientCredentialsGrantAccessTokenLifespan")),
	// 	ClientName: hc.PtrString(r.Form.Get("ClientName")),
	// 	ClientUri: hc.PtrString(r.Form.Get("ClientUri")),
	// 	Contacts: strings.Split(r.Form.Get("Contacts"), ","),
	// 	FrontchannelLogoutSessionRequired: hc.PtrBool(r.Form.Get("FrontchannelLogoutSessionRequired") == "true"),
	// 	FrontchannelLogoutUri: hc.PtrString(r.Form.Get("FrontchannelLogoutUri")),
	// 	GrantTypes: strings.Split(r.Form.Get("GrantTypes"), ","),
	// 	ImplicitGrantAccessTokenLifespan: hc.PtrString(r.Form.Get("ImplicitGrantAccessTokenLifespan")),
	// 	ImplicitGrantIdTokenLifespan: hc.PtrString(r.Form.Get("ImplicitGrantIdTokenLifespan")),
	// 	Jwks: commons.ToInterface(r.Form.Get("Jwks")),
	// 	JwksUri: hc.PtrString(r.Form.Get("JwksUri")),
	// 	JwtBearerGrantAccessTokenLifespan: hc.PtrString(r.Form.Get("JwtBearerGrantAccessTokenLifespan")),
	// 	LogoUri: hc.PtrString(r.Form.Get("LogoUri")),
	// 	Owner: hc.PtrString(r.Form.Get("Owner")),
	// 	PolicyUri: hc.PtrString(r.Form.Get("PolicyUri")),
	// 	PostLogoutRedirectUris: strings.Split(r.Form.Get("PostLogoutRedirectUris"), ","),
	// 	RedirectUris: strings.Split(r.Form.Get("RedirectUris"), ","),
	// 	RefreshTokenGrantAccessTokenLifespan: hc.PtrString(r.Form.Get("RefreshTokenGrantAccessTokenLifespan")),
	// 	RefreshTokenGrantIdTokenLifespan: hc.PtrString(r.Form.Get("RefreshTokenGrantIdTokenLifespan")),
	// 	RefreshTokenGrantRefreshTokenLifespan: hc.PtrString(r.Form.Get("RefreshTokenGrantRefreshTokenLifespan")),
	// 	RegistrationAccessToken: hc.PtrString(r.Form.Get("RegistrationAccessToken")),
	// 	RegistrationClientUri: hc.PtrString(r.Form.Get("RegistrationClientUri")),
	// 	RequestObjectSigningAlg: hc.PtrString(r.Form.Get("RequestObjectSigningAlg")),
	// 	RequestUris: strings.Split(r.Form.Get("RequestUris"), ","),
	// 	ResponseTypes: strings.Split(r.Form.Get("ResponseTypes"), ","),
	// 	Scope: hc.PtrString(r.Form.Get("Scope")),
	// 	SectorIdentifierUri: hc.PtrString(r.Form.Get("SectorIdentifierUri")),
	// 	SkipConsent: hc.PtrBool(r.Form.Get("SkipConsent") == "true"),
	// 	SkipLogoutConsent: hc.PtrBool(r.Form.Get("SkipLogoutConsent") == "true"),
	// 	SubjectType: hc.PtrString(r.Form.Get("SubjectType")),
	// 	TokenEndpointAuthMethod: hc.PtrString(r.Form.Get("TokenEndpointAuthMethod")),
	// 	TokenEndpointAuthSigningAlg: hc.PtrString(r.Form.Get("TokenEndpointAuthSigningAlg")),
	// 	TosUri: hc.PtrString(r.Form.Get("TosUri")),
	// 	UserinfoSignedResponseAlg: hc.PtrString(r.Form.Get("UserinfoSignedResponseAlg")),
	// }

	// Save the changes
	// Display new client details
	// fmt.Printf("Client Details:\n")
	// fmt.Printf("New client details : %v\n", client)

	// Update the client
	err := hydra.UpdateOAuth2ClientUsingJsonPatch(clientID, r.Form)
	if err != nil {
		logger.Logger.Error("Error updating the client", "error", err)
		http.Error(w, "Error updating the client", http.StatusInternalServerError)
		return
	}

	// Redirect to the index page and display a success message
	http.Redirect(w, r, "/", http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`<div class="notification is-success">Client %s with ID %s updated successfully</div>`, r.Form.Get("ClientName"), clientID)))
}
