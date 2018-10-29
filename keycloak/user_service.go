package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// UserService interacts with all user resources
type UserService service

// NewUserService returns a new user service for working with user resources
// in a realm.
func NewUserService(c *Client) *UserService {
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
func (us *UserService) Find(ctx context.Context, realm string, params map[string]string) ([]UserRepresentation, error) {

	path := "/realms/{realm}/users"

	var user []UserRepresentation

	_, err := us.client.newRequest(ctx).
		SetQueryParams(params).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&user).
		Get(path)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create creates a new user and returns the ID
// Response is a 201 with a location redirect
func (us *UserService) Create(ctx context.Context, realm string, user *UserRepresentation) (string, error) {
	path := "/realms/{realm}/users"

	response, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
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
func (us *UserService) Get(ctx context.Context, realm string, userID string) (*UserRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}"

	user := &UserRepresentation{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(user).
		Get(path)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Get returns a user in a realm
func (us *UserService) GetAll(ctx context.Context, realm string) ([]UserRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/users"

	var users []UserRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(users).
		Get(path)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update user information
// Response is a 204: No Content
func (us *UserService) Update(ctx context.Context, realm string, user *UserRepresentation) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    user.ID,
		}).
		SetBody(user).
		Put(path)

	return err

}

// Delete user information
// Response is a 204: No Content
func (us *UserService) Delete(ctx context.Context, realm string, userID string) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Delete(path)

	return err
}

// Impersonate user
func (us *UserService) Impersonate(ctx context.Context, realm string, userID string) (AttributeMap, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/impersonation"

	a := AttributeMap{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&a).
		Post(path)

	return a, err
}

// Count gets user count in a realm
func (us *UserService) Count(ctx context.Context, realm string) (uint32, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/count"

	var result uint32

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&result).
		Get(path)

	return result, err
}

// GetGroups gets the groups a realm user belongs to
func (us *UserService) GetGroups(ctx context.Context, realm string, userID string) ([]GroupRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/groups"

	var groups []GroupRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&groups).
		Get(path)

	return groups, err
}

// GetConsents gets consents granted by the user
func (us *UserService) GetConsents(ctx context.Context, realm string, userID string) (AttributeMap, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/consents"

	var consents AttributeMap

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&consents).
		Get(path)

	return consents, err
}

// RevokeClientConsents revokes consent and offline tokens for particular client from user
func (us *UserService) RevokeClientConsents(ctx context.Context, realm string, userID string, clientID string) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/consents/{client}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":  realm,
			"id":     userID,
			"client": clientID,
		}).
		Delete(path)

	return err
}

// DisableCredentials disables credentials of certain types for a user
func (us *UserService) DisableCredentials(ctx context.Context, realm string, userID string, credentialTypes []string) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/disable-credential-types"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Put(path)

	return err
}

// AddGroup adds a user to a group
func (us *UserService) AddGroup(ctx context.Context, realm string, userID string, groupID string) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/groups/{groupId}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":   realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Put(path)

	return err
}

// RemoveGroup removes a user from a group
func (us *UserService) RemoveGroup(ctx context.Context, realm string, userID string, groupID string) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/groups/{groupId}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":   realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Delete(path)

	return err
}

// Logout revokes all user sessions
func (us *UserService) Logout(ctx context.Context, realm string, userID string) error {

	path := "/realms/{realm}/users/{id}/logout"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Post(path)

	return err
}

// GetSessions for user
func (us *UserService) GetSessions(ctx context.Context, realm string, userID string) ([]UserSessionRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/sessions"

	var sessions []UserSessionRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&sessions).
		Get(path)

	return sessions, err
}

// GetOfflineSessions for particular client and user
func (us *UserService) GetOfflineSessions(ctx context.Context, realm string, userID string, clientID string) ([]UserSessionRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/offline-sessions/{clientId}"

	var sessions []UserSessionRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":  realm,
			"id":     userID,
			"client": clientID,
		}).
		SetResult(&sessions).
		Get(path)

	return sessions, err
}

// ResetPassword for user
func (us *UserService) ResetPassword(ctx context.Context, realm string, userID string, tempPassword *CredentialRepresentation) error {

	// nolint: goconst
	path := "/realms/{realm}/users/{id}/reset-password"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetBody(tempPassword).
		Put(path)

	return err
}
