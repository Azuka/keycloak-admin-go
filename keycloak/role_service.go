package keycloak

import "context"

// UserService interacts with all user resources
type RoleService service

// NewUserService returns a new user service for working with user resources
// in a realm.
func NewRoleService(c *Client) *RoleService {
	return &RoleService{
		client: c,
	}
}

// Get returns a user in a realm
func (us *RoleService) GetAllMapping(ctx context.Context, realm string, groupId string) (*MappingRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/groups/{id}/role-mappings"

	mappings := &MappingRepresentation{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    groupId,
		}).
		SetResult(mappings).
		Get(path)

	if err != nil {
		return nil, err
	}

	return mappings, nil
}

// Create a new role in realm
func (us *RoleService) CreateRole(ctx context.Context, realm string, roleName string, description string) error {

	// nolint: goconst
	path := "/realms/{realm}/roles"

	role := RoleRepresentation{Name: roleName, Description: description}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(&role).
		Post(path)

	return err
}

// Get returns a user in a realm
func (us *RoleService) GetAll(ctx context.Context, realm string) ([]RoleRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/roles"

	var roles []RoleRepresentation

	_, err := us.client.newRequest(ctx).
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
