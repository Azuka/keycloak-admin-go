package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// ClientsService interacts with all user resources
type ClientsService service

// ClientsS returns a new client service for working with client resources
// in a realm.
func (c *Client) Clients() *ClientsService {
	return &ClientsService{
		client: c,
	}
}

// Create creates a new client and returns the ID
// Response is a 201 with a location redirect
func (s *ClientsService) Create(ctx context.Context, realm string, client *ClientRepresentation) (string, error) {
	path := "/realms/{realm}/clients"

	response, err := s.client.newRequest(ctx).
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
func (s *ClientsService) List(ctx context.Context, realm string) ([]ClientRepresentation, error) {

	path := "/realms/{realm}/clients"

	var clients []ClientRepresentation

	_, err := s.client.newRequest(ctx).
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
