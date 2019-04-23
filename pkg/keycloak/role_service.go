package keycloak

import (
	"net/url"
	"strings"
)

// RoleService interacts with all role resources
type RoleService service

// Roles returns a new role service for working with role resources
// in a realm.
func (c *Client) Roles() *RoleService {
	return &RoleService{
		client: c,
	}
}

// Create creates a new role in realm
func (s *RoleService) Create(realm string, role *RoleRepresentation) (string, error) {

	path := "/realms/{realm}/roles"

	response, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(role).
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

// Get gets a roles by name returns all roles
func (s *RoleService) Get(realm string) ([]RoleRepresentation, error) {

	path := "/realms/{realm}/roles"

	var roles []RoleRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&roles).
		Get(path)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

// List returns all roles
func (s *RoleService) List(realm string) ([]RoleRepresentation, error) {

	path := "/realms/{realm}/roles"

	var roles []RoleRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&roles).
		Get(path)

	if err != nil {
		return nil, err
	}

	return roles, nil
}
