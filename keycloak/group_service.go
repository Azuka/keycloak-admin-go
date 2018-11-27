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
func (us *GroupService) Create(ctx context.Context, realm string, groupName string) error {

	// nolint: goconst
	path := "/realms/{realm}/groups"
	group := &GroupRepresentation{}
	group.Name = groupName

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(group).
		Post(path)

	if err != nil {
		return err
	}

	return nil
}

// Get returns a user in a realm
func (us *GroupService) Delete(ctx context.Context, realm string, groupId string) error {

	// nolint: goconst
	path := "/realms/{realm}/groups/{groupId}"
	group := &GroupRepresentation{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":   realm,
			"groupId": groupId,
		}).
		SetBody(group).
		Delete(path)

	if err != nil {
		return err
	}

	return nil
}

// Get returns a user in a realm
func (us *GroupService) Get(ctx context.Context, realm string, groupId string) (*GroupRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/groups/{id}"
	group := &GroupRepresentation{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    groupId,
		}).
		SetResult(group).
		Get(path)

	if err != nil {
		return nil, err
	}

	return group, nil
}

// Get returns a user in a realm
func (us *GroupService) AddRole(ctx context.Context, realm string, groupId string, role RoleRepresentation) error {

	// nolint: goconst
	path := "/realms/{realm}/groups/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    groupId,
		}).
		SetBody(roles).
		Post(path)

	return err
}

// Get returns a user in a realm
func (us *GroupService) DeleteRole(ctx context.Context, realm string, groupId string, role RoleRepresentation) error {

	// nolint: goconst
	path := "/realms/{realm}/groups/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    groupId,
		}).
		SetBody(roles).
		Delete(path)

	return err
}

// Get returns a user in a realm
func (us *GroupService) GetAllDetail(ctx context.Context, realm string) ([]GroupRepresentation, error) {

	var result []GroupRepresentation

	groups, err := us.GetAll(ctx, realm)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		g, err := us.Get(ctx, realm, group.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *g)
	}

	return result, nil
}

// Get returns a user in a realm
func (us *GroupService) GetAll(ctx context.Context, realm string) ([]GroupRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/groups"
	var groups []GroupRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&groups).
		Get(path)

	if err != nil {
		return nil, err
	}

	return groups, nil
}
