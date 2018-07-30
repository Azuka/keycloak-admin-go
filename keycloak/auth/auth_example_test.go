package auth_test

import (
	"context"
	"github.com/Azuka/keycloak-admin-go/keycloak/auth"
	"net/url"
)

func ExampleConfig_Client() {
	config := auth.Config{
		ClientID: "admin-cli",
		TokenURL: "https://keycloak.local/auth/realms/master/protocol/openid-connect/token",
		EndpointParams: url.Values{
			"username":   {"keycloak"},
			"password":   {"changeme"},
			"grant_type": {"password"},
			"client_id":  {"admin-cli"},
		},
	}

	client := config.Client(context.Background())

	// This will make an authenticated request
	client.Get("https://keycloak.local/auth/admin/realms/master/users?username=keycloak-admin")
}
