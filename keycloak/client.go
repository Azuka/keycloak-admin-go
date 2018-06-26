// Package keycloak contains a client and relevant data structs for interacting
// with the Keycloak Admin REST API: https://www.keycloak.org/docs-api/4.0/rest-api/index.html
package keycloak

import (
	"net/http"
	"net/url"

	"strings"

	"context"
	"fmt"
	"github.com/go-resty/resty"
)

const userAgent = "go/keycloak-admin"

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
	restClient.SetRedirectPolicy(resty.DomainCheckRedirectPolicy(u.Hostname()))

	// Strip out trailing slash
	u.Path = strings.TrimRight(u.Path, "/")

	client := &Client{
		BaseURL:    u,
		restClient: restClient,
	}

	client.Users = NewUserService(client)

	return client
}

// Debug enables debugging for requests
func (c *Client) Debug() {
	c.restClient.SetDebug(true)
}

// newRequest creates a new request
func (c *Client) newRequest(ctx context.Context) *resty.Request {
	return c.restClient.
		R().
		//SetError(&Error{}).
		SetContext(ctx).
		SetHeader("UserAgent", userAgent)
}

func (c *Client) exec(response *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if response.StatusCode() < 400 {
		return nil
	}

	return &Error{
		Message: fmt.Sprintf("%s %s: %s", response.Request.Method, response.Request.URL, response.Status()),
		Code:    response.StatusCode(),
	}

}
