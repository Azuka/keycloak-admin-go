package auth_test

import (
	"context"

	"github.com/Azuka/keycloak-admin-go/keycloak/auth"
)

func ExampleConfig_Client() {
	config := auth.Config{
		ClientID:  "admin-cli",
		TokenURL:  "https://keycloak.local/auth/realms/master/protocol/openid-connect/token",
		Username:  "keycloak",
		Password:  "password",
		GrantType: auth.PasswordGrant,
	}

	client := config.Client(context.Background())

	// This will make an authenticated request
	_, _ = client.Get("https://keycloak.local/auth/admin/realms/master/users?username=keycloak-admin")
}

func ExampleConfig_Client_client_credentials() {
	config := auth.Config{
		ClientID:     "admin-cli",
		TokenURL:     "https://keycloak.local/auth/realms/master/protocol/openid-connect/token",
		ClientSecret: "my-secret",
		GrantType:    auth.ClientCredentialsGrant,
	}

	client := config.Client(context.Background())

	// This will make an authenticated request
	_, _ = client.Get("https://keycloak.local/auth/admin/realms/master/users?username=keycloak-admin")
}
