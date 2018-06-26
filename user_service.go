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

// Create creates a new user and returns the ID
func (us *UserService) Create(ctx context.Context, realm string, user *UserRepresentation) (string, error) {
	u := us.client.BaseURL
	u.Path += "/{realm}/users"

	response, err := us.client.newRequest().
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(user).
		SetContext(ctx).
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
	u.Path += "/{realm}/users/{userID}"

	user := &UserRepresentation{}

	_, err := us.client.newRequest().
		SetPathParams(map[string]string{
			"realm":  realm,
			"userID": userID,
		}).
		SetContext(ctx).
		SetResult(user).
		Get(u.String())

	return user, err
}
