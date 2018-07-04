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

	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users"

	var user []UserRepresentation

	_, err := us.client.newRequest(ctx).
		SetQueryParams(params).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&user).
		Get(u.String())

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create creates a new user and returns the ID
// Response is a 201 with a location redirect
func (us *UserService) Create(ctx context.Context, realm string, user *UserRepresentation) (string, error) {
	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users"

	response, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(user).
		SetResult(user).
		Post(u.String())

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

	u := us.client.BaseURL
	// nolint: goconst
	u.Path += "/realms/{realm}/users/{id}"

	user := &UserRepresentation{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(user).
		Get(u.String())

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update user information
// Response is a 204: No Content
func (us *UserService) Update(ctx context.Context, realm string, user *UserRepresentation) error {

	u := us.client.BaseURL
	// nolint: goconst
	u.Path += "/realms/{realm}/users/{id}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    user.ID,
		}).
		SetBody(user).
		Put(u.String())

	return err

}

// Delete user information
// Response is a 204: No Content
func (us *UserService) Delete(ctx context.Context, realm string, userID string) error {

	u := us.client.BaseURL
	// nolint: goconst
	u.Path += "/realms/{realm}/users/{id}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Delete(u.String())

	return err
}

// Impersonate user
func (us *UserService) Impersonate(ctx context.Context, realm string, userID string) (AttributeMap, error) {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/impersonation"

	a := AttributeMap{}

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&a).
		Post(u.String())

	return a, err
}

// Count gets user count in a realm
func (us *UserService) Count(ctx context.Context, realm string) (uint32, error) {

	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users/count"

	var result uint32

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&result).
		Get(u.String())

	return result, err
}

// GetGroups gets the groups a realm user belongs to
func (us *UserService) GetGroups(ctx context.Context, realm string, userID string) ([]GroupRepresentation, error) {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/groups"

	var groups []GroupRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&groups).
		Get(u.String())

	return groups, err
}

// GetConsents gets consents granted by the user
func (us *UserService) GetConsents(ctx context.Context, realm string, userID string) (AttributeMap, error) {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/consents"

	var consents AttributeMap

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&consents).
		Get(u.String())

	return consents, err
}

// RevokeClientConsents revokes consent and offline tokens for particular client from user
func (us *UserService) RevokeClientConsents(ctx context.Context, realm string, userID string, clientID string) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/consents/{client}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":  realm,
			"id":     userID,
			"client": clientID,
		}).
		Delete(u.String())

	return err
}

// DisableCredentials disables credentials of certain types for a user
func (us *UserService) DisableCredentials(ctx context.Context, realm string, userID string, credentialTypes []string) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/disable-credential-types"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Put(u.String())

	return err
}

// AddGroup adds a user to a group
func (us *UserService) AddGroup(ctx context.Context, realm string, userID string, groupID string) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/groups/{groupId}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":   realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Put(u.String())

	return err
}

// RemoveGroup removes a user from a group
func (us *UserService) RemoveGroup(ctx context.Context, realm string, userID string, groupID string) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/groups/{groupId}"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":   realm,
			"id":      userID,
			"groupId": groupID,
		}).
		Delete(u.String())

	return err
}

// Logout revokes all user sessions
func (us *UserService) Logout(ctx context.Context, realm string, userID string) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/logout"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		Post(u.String())

	return err
}

// GetSessions for user
func (us *UserService) GetSessions(ctx context.Context, realm string, userID string) ([]UserSessionRepresentation, error) {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/sessions"

	var sessions []UserSessionRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetResult(&sessions).
		Get(u.String())

	return sessions, err
}

// GetOfflineSessions for particular client and user
func (us *UserService) GetOfflineSessions(ctx context.Context, realm string, userID string, clientID string) ([]UserSessionRepresentation, error) {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/offline-sessions/{clientId}"

	var sessions []UserSessionRepresentation

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":  realm,
			"id":     userID,
			"client": clientID,
		}).
		SetResult(&sessions).
		Get(u.String())

	return sessions, err
}

// ResetPassword for user
func (us *UserService) ResetPassword(ctx context.Context, realm string, userID string, tempPassword *CredentialRepresentation) error {

	u := us.client.BaseURL
	u.Path += "/{realm}/users/{id}/reset-password"

	_, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
			"id":    userID,
		}).
		SetBody(tempPassword).
		Put(u.String())

	return err
}
