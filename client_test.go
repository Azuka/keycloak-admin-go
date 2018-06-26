package keycloak_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/azuka/keycloak-admin-go"
	"github.com/stretchr/testify/assert"
)

func ExampleNewClient() {
	u, _ := url.Parse("http://localhost/auth/admin")
	c := keycloak.NewClient(*u, http.DefaultClient)
	userID, _ := c.Users.Create(context.Background(), "myRealm", &keycloak.UserRepresentation{
		Username: "hello-world",
	})
	fmt.Println("UserID: ", userID)
	// Output: UserID:
}

func TestNewClient(t *testing.T) {
	a := assert.New(t)

	url, _ := url.Parse("http://localhost/keycloak/auth/admin/")

	client := keycloak.NewClient(*url, http.DefaultClient)

	a.NotNil(client)
	a.NotNil(client)
	a.NotNil(client.Users)
	a.Equal("/keycloak/auth/admin", client.BaseURL.Path)
}
