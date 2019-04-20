package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// ClientsService interacts with all user resources
type ClientsService service

// NewClientsService returns a new user service for working with user resources
// in a realm.
func NewClientsService(c *Client) *ClientsService {
	return &ClientsService{
		client: c,
	}
}

// Create creates a new client and returns the ID
// Response is a 201 with a location redirect
func (cs *ClientsService) Create(ctx context.Context, realm string, client *ClientRepresentation) (string, error) {
	path := "/realms/{realm}/clients"

	response, err := cs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(client).
		Post(path)

	if err != nil {
		return "", err
	}

	location, err := url.Parse(response.Header().Get("Location"))

	if err != nil {
		return "", err
	}

	components := strings.Split(location.Path, "/")

	return components[len(components)-1], nil
}

// List clients
func (cs *ClientsService) List(ctx context.Context, realm string) ([]ClientRepresentation, error) {

	path := "/realms/{realm}/clients"

	var clients []ClientRepresentation

	_, err := cs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&clients).
		Get(path)

	if err != nil {
		return nil, err
	}

	return clients, nil
}
