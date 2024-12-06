package hydra

import (
	"context"
	"orydra/config"

	hc "github.com/ory/hydra-client-go/v2"
)

var (
	hydraConfig *hc.Configuration
	HydraClient hc.OAuth2Client
)

func init() {
	envVars := config.SetEnv()

	// Hydra API Client Configuration
	hydraConfig = hc.NewConfiguration()
	hydraConfig.Servers = []hc.ServerConfiguration{
		{
			URL: envVars.HYDRA_ADMIN_URL,
		},
	}
}

func CreateNewHydraClient(ctx context.Context, oAuth2Client *hc.OAuth2Client) (*hc.OAuth2Client, error) {
	// Create a new OAuth2 client with oAuth2Client as the request body
	client, _, err := hc.NewAPIClient(hydraConfig).OAuth2API.CreateOAuth2Client(ctx).OAuth2Client(*oAuth2Client).Execute()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetHydraClient(ctx context.Context, clientID string) (*hc.OAuth2Client, error) {
	client, _, err := hc.NewAPIClient(hydraConfig).OAuth2API.GetOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func DeleteHydraClient(ctx context.Context, clientID string) error {
	_, err := hc.NewAPIClient(hydraConfig).OAuth2API.DeleteOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		return err
	}

	return nil
}

// Get the list of all clients
func GetAllHydraClients(ctx context.Context) ([]hc.OAuth2Client, error) {
	clients, _, err := hc.NewAPIClient(hydraConfig).OAuth2API.ListOAuth2Clients(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return clients, nil
}

// Get specific client details by ID
func GetHydraClientByID(ctx context.Context, clientID string) (*hc.OAuth2Client, error) {
	client, _, err := hc.NewAPIClient(hydraConfig).OAuth2API.GetOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Get clientID by client name
func GetHydraClientIDByName(ctx context.Context, clientName string) (string, error) {
	clients, err := GetAllHydraClients(ctx)
	if err != nil {
		return "", err
	}

	for _, client := range clients {
		if *client.ClientName == clientName {
			return *client.ClientId, nil
		}
	}

	return "", nil
}

// Get specific client details by client name
func GetHydraClientByName(ctx context.Context, clientName string) (*hc.OAuth2Client, error) {
	clientID, err := GetHydraClientIDByName(ctx, clientName)
	if err != nil {
		return nil, err
	}

	client, _, err := hc.NewAPIClient(hydraConfig).OAuth2API.GetOAuth2Client(ctx, clientID).Execute()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Get client secret by client ID
func GetHydraClientSecretByID(ctx context.Context, clientID string) (string, error) {
	client, err := GetHydraClientByID(ctx, clientID)
	if err != nil {
		return "", err
	}

	clientSecret := client.GetClientSecret()

	return clientSecret, nil
}
