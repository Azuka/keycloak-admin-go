package keycloak

import (
	"context"
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
func (s *RoleService) Create(ctx context.Context, role *RoleRepresentation) (string, error) {

	path := "/realms/{realm}/roles"

	response, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
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

// Delete deletes the role in realm
func (s *RoleService) Delete(ctx context.Context, role *RoleRepresentation) error {

	path := "/realms/{realm}/roles/{role-name}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"role-name": role.Name,
		}).
		SetBody(role).
		Delete(path)

	return err
}

// Get gets a role by name
func (s *RoleService) Get(ctx context.Context, name string) (*RoleRepresentation, error) {

	path := "/realms/{realm}/roles/{role-name}"

	role := &RoleRepresentation{}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"role-name": name,
		}).
		SetResult(role).
		Get(path)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// List returns all roles
func (s *RoleService) List(ctx context.Context) ([]RoleRepresentation, error) {

	path := "/realms/{realm}/roles"

	var roles []RoleRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetResult(&roles).
		Get(path)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

// GetComposites gets the composite roles from the given role
func (s *RoleService) GetComposites(ctx context.Context, role *RoleRepresentation) ([]RoleRepresentation, error) {
	path := "/realms/{realm}/roles/{role-name}/composites"

	var roles []RoleRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"role-name": role.Name,
		}).
		SetResult(&roles).
		Get(path)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

// AddComposite adds composite roles to role
func (s *RoleService) AddComposite(ctx context.Context, role *RoleRepresentation, composites []RoleRepresentation) error {
	path := "/realms/{realm}/roles/{role-name}/composites"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"role-name": role.Name,
		}).
		SetBody(&composites).
		Post(path)

	return err
}

// RemoveComposite removes composite roles from role
func (s *RoleService) RemoveComposite(ctx context.Context, role *RoleRepresentation, composites []RoleRepresentation) error {
	path := "/realms/{realm}/roles/{role-name}/composites"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"role-name": role.Name,
		}).
		SetBody(&composites).
		Delete(path)

	return err
}
