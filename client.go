// Package keycloak contains a client and relevant data structs for interacting
// with the Keycloak Admin REST API: https://www.keycloak.org/docs-api/4.0/rest-api/index.html
package keycloak

import (
	"net/http"
	"net/url"

	"strings"

	"github.com/go-resty/resty"
)

// Client is the API client for talking to keycloak admin
type Client struct {
	BaseURL    url.URL
	restClient *resty.Client

	// Services for working with various keycloak resources
	Users *UserService
}

// NewClient creates a new client instance set to talk to the keycloak service
// as well as the various services for working with specific resources
func NewClient(u url.URL, c *http.Client) *Client {

	restClient := resty.NewWithClient(c)

	// Strip out trailing slash
	u.Path = strings.TrimRight(u.Path, "/")

	client := &Client{
		BaseURL:    u,
		restClient: restClient,
	}

	client.Users = NewUserService(client)

	return client
}

// newRequest creates a new request
func (c *Client) newRequest() *resty.Request {
	return c.restClient.R()
}
