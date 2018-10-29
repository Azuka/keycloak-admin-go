package keycloak

import (
	"context"
)

// UserService interacts with all user resources
type GroupService service

// NewUserService returns a new user service for working with user resources
// in a realm.
func NewGroupService(c *Client) *GroupService {
	return &GroupService{
		client: c,
	}
}

// Get returns a user in a realm
func (us *GroupService) GetAll(ctx context.Context, realm string) ([]GroupRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/groups"
	var groups []GroupRepresentation

	_, err := us.client.newRequest(ctx).
		SetResult(&groups).
		Get(path)

	if err != nil {
		return nil, err
	}

	return groups, nil
}
