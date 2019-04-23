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

// Create creates a new role in realm
func (s *RoleService) Create(ctx context.Context, roleName string, description string) error {

	path := "/realms/{realm}/roles"

	role := RoleRepresentation{Name: roleName, Description: description}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetBody(&role).
		Post(path)

	return err
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
