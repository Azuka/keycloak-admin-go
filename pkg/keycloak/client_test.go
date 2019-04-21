package keycloak

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testRelam = "test-realm"

func ExampleNewClient() {
	u, _ := url.Parse("http://localhost/auth/admin")
	c := NewClient(*u, http.DefaultClient, testRelam)
	userID, _ := c.Users().Create(context.Background(), &UserRepresentation{
		Username: "hello-world",
	})
	fmt.Println("UserID: ", userID)
}

func TestNewClient(t *testing.T) {
	a := assert.New(t)

	url, _ := url.Parse("http://localhost/keycloak/auth/admin/")

	client := NewClient(*url, http.DefaultClient, testRelam)

	a.NotNil(client)
	a.NotNil(client)
	a.False(client.restClient.Debug)
}

func TestNewClientDebug(t *testing.T) {
	a := assert.New(t)

	client := NewClient(url.URL{}, http.DefaultClient, testRelam)
	client.Debug()
	a.True(client.restClient.Debug)
}
