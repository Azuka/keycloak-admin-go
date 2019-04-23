package keycloak

// UserService interacts with all user resources
type GroupService service

// Groups returns a new group service for working with group resources
// in a realm.
func (c *Client) Groups() *GroupService {
	return &GroupService{
		client: c,
	}
}

// Create creates a group in a realm
func (s *GroupService) Create(groupName string) error {

	path := "/realms/{realm}/groups"
	group := &GroupRepresentation{}
	group.Name = groupName

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetBody(group).
		Post(path)

	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a group from a realm
func (s *GroupService) Delete(groupID string) error {

	path := "/realms/{realm}/groups/{groupID}"
	group := &GroupRepresentation{}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm":   s.client.Realm,
			"groupID": groupID,
		}).
		SetBody(group).
		Delete(path)

	if err != nil {
		return err
	}

	return nil
}

// Get returns a group in a realm
func (s *GroupService) Get(groupID string) (*GroupRepresentation, error) {

	path := "/realms/{realm}/groups/{id}"
	group := &GroupRepresentation{}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    groupID,
		}).
		SetResult(group).
		Get(path)

	if err != nil {
		return nil, err
	}

	return group, nil
}

// AddRole adds a role to a group in a realm
func (s *GroupService) AddRole(groupID string, role RoleRepresentation) error {

	path := "/realms/{realm}/groups/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    groupID,
		}).
		SetBody(roles).
		Post(path)

	return err
}

// DeleteRole deletes a role from a group in a realm
func (s *GroupService) DeleteRole(groupID string, role RoleRepresentation) error {

	path := "/realms/{realm}/groups/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    groupID,
		}).
		SetBody(roles).
		Delete(path)

	return err
}

// ListDetail returns all groups in realm with details
func (s *GroupService) ListDetail(realm string) ([]GroupRepresentation, error) {

	var result []GroupRepresentation

	groups, err := s.List(realm)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		g, err := s.Get(group.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *g)
	}

	return result, nil
}

// List returns all groups in a realm
func (s *GroupService) List(realm string) ([]GroupRepresentation, error) {

	path := "/realms/{realm}/groups"
	var groups []GroupRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetResult(&groups).
		Get(path)

	if err != nil {
		return nil, err
	}

	return groups, nil
}

// ListMapping returns a all role mappings for group (TODO: maybe put this as a GroupRepresentation Method)
func (s *GroupService) ListMapping(groupID string) (*MappingRepresentation, error) {

	path := "/realms/{realm}/groups/{id}/role-mappings"

	mappings := &MappingRepresentation{}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    groupID,
		}).
		SetResult(mappings).
		Get(path)

	if err != nil {
		return nil, err
	}

	return mappings, nil
}
