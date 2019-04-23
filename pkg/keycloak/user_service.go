package keycloak

import (
	"net/url"
	"strings"
)

// UserService interacts with all user resources
type UserService service

// Users returns a new user service for working with user resources
// in a realm.
func (c *Client) Users() *UserService {
	return &UserService{
		client: c,
	}
}

// Find returns users based on query params
// Params:
// - email
// - first
// - firstName
// - lastName
// - max
// - search
// - userName
func (s *UserService) Find(params map[string]string) ([]UserRepresentation, error) {

	path := "/realms/{realm}/users"

	var users []UserRepresentation

	_, err := s.client.newRequest().
		SetQueryParams(params).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetResult(&users).
		Get(path)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create creates a new user and returns the ID
// Response is a 201 with a location redirect
func (s *UserService) Create(user *UserRepresentation) (string, error) {

	path := "/realms/{realm}/users"

	response, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetBody(user).
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

// Get returns a user in a realm
func (s *UserService) Get(userID string) (*UserRepresentation, error) {

	path := "/realms/{realm}/users/{id}"

	user := &UserRepresentation{}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetResult(user).
		Get(path)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// List returns a user in a realm
func (s *UserService) List(realm string) ([]UserRepresentation, error) {

	path := "/realms/{realm}/users"

	var users []UserRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetResult(&users).
		Get(path)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update user information
// Response is a 204: No Content
func (s *UserService) Update(user *UserRepresentation) error {

	path := "/realms/{realm}/users/{id}"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    user.ID,
		}).
		SetBody(user).
		Put(path)

	return err

}

// Delete user information
// Response is a 204: No Content
func (s *UserService) Delete(userID string) error {

	path := "/realms/{realm}/users/{id}"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		Delete(path)

	return err
}

// Impersonate user
func (s *UserService) Impersonate(userID string) (AttributeMap, error) {

	path := "/realms/{realm}/users/{id}/impersonation"

	a := AttributeMap{}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetResult(&a).
		Post(path)

	return a, err
}

// Count gets user count in a realm
func (s *UserService) Count(realm string) (uint32, error) {

	path := "/realms/{realm}/users/count"

	var result uint32

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetResult(&result).
		Get(path)

	return result, err
}

// GetGroups gets the groups a realm user belongs to
func (s *UserService) GetGroups(userID string) ([]GroupRepresentation, error) {

	path := "/realms/{realm}/users/{id}/groups"

	var groups []GroupRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetResult(&groups).
		Get(path)

	return groups, err
}

// GetConsents gets consents granted by the user
func (s *UserService) GetConsents(userID string) (AttributeMap, error) {

	path := "/realms/{realm}/users/{id}/consents"

	var consents AttributeMap

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetResult(&consents).
		Get(path)

	return consents, err
}

// RevokeClientConsents revokes consent and offline tokens for particular client from user
func (s *UserService) RevokeClientConsents(userID string, clientID string) error {

	path := "/realms/{realm}/users/{id}/consents/{client}"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm":  s.client.Realm,
			"id":     userID,
			"client": clientID,
		}).
		Delete(path)

	return err
}

// DisableCredentials disables credentials of certain types for a user
func (s *UserService) DisableCredentials(userID string, credentialTypes []string) error {

	path := "/realms/{realm}/users/{id}/disable-credential-types"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		Put(path)

	return err
}

// AddGroup adds a user to a group
func (s *UserService) AddGroup(userID string, groupID string) error {

	path := "/realms/{realm}/users/{id}/groups/{groupId}"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm":   s.client.Realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Put(path)

	return err
}

// RemoveGroup removes a user from a group
func (s *UserService) RemoveGroup(userID string, groupID string) error {

	path := "/realms/{realm}/users/{id}/groups/{groupId}"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm":   s.client.Realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Delete(path)

	return err
}

// Logout revokes all user sessions
func (s *UserService) Logout(userID string) error {

	path := "/realms/{realm}/users/{id}/logout"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		Post(path)

	return err
}

// GetSessions for user
func (s *UserService) GetSessions(userID string) ([]UserSessionRepresentation, error) {

	path := "/realms/{realm}/users/{id}/sessions"

	var sessions []UserSessionRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetResult(&sessions).
		Get(path)

	return sessions, err
}

// GetOfflineSessions for particular client and user
func (s *UserService) GetOfflineSessions(userID string, clientID string) ([]UserSessionRepresentation, error) {

	path := "/realms/{realm}/users/{id}/offline-sessions/{clientId}"

	var sessions []UserSessionRepresentation

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm":  s.client.Realm,
			"id":     userID,
			"client": clientID,
		}).
		SetResult(&sessions).
		Get(path)

	return sessions, err
}

// ResetPassword for user
func (s *UserService) ResetPassword(userID string, tempPassword *CredentialRepresentation) error {

	path := "/realms/{realm}/users/{id}/reset-password"

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetBody(tempPassword).
		Put(path)

	return err
}

// AddRole adds a role to a user in a realm
func (s *UserService) AddRole(userID string, role RoleRepresentation) error {

	path := "/realms/{realm}/users/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetBody(roles).
		Post(path)

	return err
}

// DeleteRole deletes a role from a user in a realm
func (s *UserService) DeleteRole(userID string, role RoleRepresentation) error {
	path := "/realms/{realm}/users/{id}/role-mappings/realm"
	roles := &[]RoleRepresentation{role}

	_, err := s.client.newRequest().
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    userID,
		}).
		SetBody(roles).
		Delete(path)

	return err
}
