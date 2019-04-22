package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// ClientsService interacts with all user resources
type ClientService service

// ClientsS returns a new client service for working with client resources
// in a realm.
func (c *Client) Clients() *ClientService {
	return &ClientService{
		client: c,
	}
}

// Create creates a new client and returns the ID
// Response is a 201 with a location redirect
func (s *ClientService) Create(ctx context.Context, client *ClientRepresentation) (string, error) {
	path := "/realms/{realm}/clients"

	response, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
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

// Get returns a client in a realm
func (s *ClientService) Get(ctx context.Context, ID string) (*ClientRepresentation, error) {

	path := "/realms/{realm}/clients/{id}"

	client := &ClientRepresentation{}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    ID,
		}).
		SetResult(client).
		Get(path)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Find returns clients based on query params
// Params:
// - clientId
func (s *ClientService) Find(ctx context.Context, params map[string]string) ([]ClientRepresentation, error) {

	path := "/realms/{realm}/clients"

	var clients []ClientRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetQueryParams(params).
		SetResult(&clients).
		Get(path)

	if err != nil {
		return nil, err
	}

	return clients, nil
}
