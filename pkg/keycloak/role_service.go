package keycloak

import "context"

// RoleService interacts with all role resources
type RoleService service

// Roles returns a new role service for working with role resources
// in a realm.
func (c *Client) Roles() *RoleService {
	return &RoleService{
		client: c,
	}
}

// Creates a new role in realm
func (s *RoleService) CreateRole(ctx context.Context, realm string, roleName string, description string) error {

	path := "/realms/{realm}/roles"

	role := RoleRepresentation{Name: roleName, Description: description}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(&role).
		Post(path)

	return err
}

// List returns all roles
func (s *RoleService) List(ctx context.Context, realm string) ([]RoleRepresentation, error) {

	path := "/realms/{realm}/roles"

	var roles []RoleRepresentation

	_, err := s.client.newRequest(ctx).
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
