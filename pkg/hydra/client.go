package hydra

import (
	"context"
	"fmt"
	"net/url"
	"orydra/config"
	"orydra/pkg/commons"
	"orydra/pkg/logger"
	"strings"

	hc "github.com/ory/hydra-client-go/v2"
)

type StringSliceJSONFormat []string

var (
	HydraConfig *hc.Configuration
	HydraClient hc.OAuth2Client
)

func init() {
	envVars := config.SetEnv()

	// Hydra API Client Configuration
	HydraConfig = hc.NewConfiguration()
	HydraConfig.Servers = []hc.ServerConfiguration{
		{
			URL: envVars.HYDRA_ADMIN_URL,
		},
	}
}

func CreateNewHydraClient(ctx context.Context, oAuth2Client *hc.OAuth2Client) (*hc.OAuth2Client, error) {
	// Create a new OAuth2 client with oAuth2Client as the request body
	client, _, err := hc.NewAPIClient(HydraConfig).OAuth2API.CreateOAuth2Client(ctx).OAuth2Client(*oAuth2Client).Execute()
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la création du client", "error", err, "function", funcName)
		return nil, err
	}

	return client, nil
}

func DeleteHydraClient(ctx context.Context, clientID string) error {
	_, err := hc.NewAPIClient(HydraConfig).OAuth2API.DeleteOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la suppression du client", "error", err, "function", funcName)
		return err
	}

	return nil
}

// Get the list of all clients
// @param ctx context.Context
// @return []hc.OAuth2Client, error
func GetAllHydraClients(ctx context.Context) ([]hc.OAuth2Client, error) {
	clients, _, err := hc.NewAPIClient(HydraConfig).OAuth2API.ListOAuth2Clients(ctx).Execute()
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la récupération de la liste des clients", "error", err, "function", funcName)
		return nil, err
	}

	return clients, nil
}

// Get specific client details by ID
func GetHydraClientByID(ctx context.Context, clientID string) (*hc.OAuth2Client, error) {
	client, _, err := hc.NewAPIClient(HydraConfig).OAuth2API.GetOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la récupération du client", "error", err, "function", funcName)
		return nil, err
	}

	return client, nil
}

// Get clientID by client name
func GetHydraClientIDByName(ctx context.Context, clientName string) (string, error) {
	clients, err := GetAllHydraClients(ctx)
	if err != nil {
		funcName := logger.GetFunctionName()
		logger.Logger.Error("Erreur lors de la récupération de la liste des clients", "error", err, "function", funcName)
		return "", err
	}

	for _, client := range clients {
		if *client.ClientName == clientName {
			return *client.ClientId, nil
		}
	}

	return "", nil
}

func UpdateOAuth2ClientUsingJsonPatch(clientID string, form url.Values) error {
	patches := []hc.JsonPatch{}

	ctx := context.Background()

	for key, values := range form {
		value := values[0]

		switch key {
		case "Jwks",
			"ClientSecretExpiresAt",
			"CreatedAt",
			"Metadata",
			"UpdatedAt",
			"AdditionalProperties":

			fmt.Printf("Key : %s, Value: %s\n", key, value)

			continue

		case "AllowedCorsOrigins",
			"Audience",
			"Contacts",
			"GrantTypes",
			"PostLogoutRedirectUris",
			"RedirectUris",
			"RequestUris",
			"ResponseTypes":
			if value == "" || value == "[]" {
				continue
			}

			// Traiter les champs []string
			key := commons.ToSnakeCase(key)

			patches = append(patches, hc.JsonPatch{
				Op:    "replace",
				Path:  "/" + key,
				Value: strings.Split(value, ","), // Convertir en []string
			})
			fmt.Printf("%v\n", patches)

		case "SkipConsent",
			"SkipLogoutConsent",
			"FrontchannelLogoutSessionRequired",
			"BackchannelLogoutSessionRequired":
			key := commons.ToSnakeCase(key)

			// Traiter les champs booléens
			patches = append(patches, hc.JsonPatch{
				Op:    "replace",
				Path:  "/" + key,
				Value: value == "on",
			})
			fmt.Printf("Key : %s, Value: %s\n", key, value)

		default:
			// Par défaut, traiter comme une chaîne
			key := commons.ToSnakeCase(key)

			if strings.Contains(key, "lifespan") {
				if value == "" {
					continue
				}
			}

			patches = append(patches, hc.JsonPatch{
				Op:    "replace",
				Path:  "/" + key,
				Value: value,
			})
			fmt.Printf("Key : %s, Value: %s\n", key, value)
		}
	}

	// Appliquer les patches via l'API Hydra
	_, _, err := hc.NewAPIClient(HydraConfig).OAuth2API.PatchOAuth2Client(ctx, clientID).JsonPatch(patches).Execute()
	return err
}
