package keycloak

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNewClient() {
	u, _ := url.Parse("http://localhost/auth/admin")
	c := NewClient(*u, http.DefaultClient)
	userID, _ := c.Users.Create(context.Background(), "myRealm", &UserRepresentation{
		Username: "hello-world",
	})
	fmt.Println("UserID: ", userID)
	// Output: UserID:
}

func TestNewClient(t *testing.T) {
	a := assert.New(t)

	url, _ := url.Parse("http://localhost/keycloak/auth/admin/")

	client := NewClient(*url, http.DefaultClient)

	a.NotNil(client)
	a.NotNil(client)
	a.NotNil(client.Users)
	a.False(client.restClient.Debug)
	a.Equal("/keycloak/auth/admin", client.BaseURL.Path)
}

func TestNewClientDebug(t *testing.T) {
	a := assert.New(t)

	client := NewClient(url.URL{}, http.DefaultClient)
	client.Debug()
	a.True(client.restClient.Debug)
}
