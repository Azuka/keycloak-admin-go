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
func (us *RoleService) GetAllMapping(ctx context.Context, realm string, groupId string) ([]MappingRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/groups/{id}/role-mappings"

	var mappings []MappingRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    groupId,
		}).
		SetResult(&mappings).
		Get(path)

	if err != nil {
		return nil, err
	}

	return mappings, nil
}

// Get returns a user in a realm
func (us *RoleService) GetAll(ctx context.Context, realm string) ([]RoleRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/clients/admin-cli/roles"

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
