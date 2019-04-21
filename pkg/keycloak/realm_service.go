package keycloak

import (
	"context"
)

// NewRealmService interacts with all realm resources
type RealmService service

// NewRealmService returns a new realm service for working with realm resources
func (c *Client) Realms() *RealmService {
	return &RealmService{
		client: c,
	}
}

// Get realm with realm name (not id!)
func (s *RealmService) Get(ctx context.Context, realm string) (*RealmRepresentation, error) {

	path := "/realms/{realm}"

	rr := &RealmRepresentation{}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(rr).
		Get(path)

	if err != nil {
		return nil, err
	}
	return rr, nil
}

// Get returns a user in a realm
func (s *RealmService) List(ctx context.Context) ([]RealmRepresentation, error) {

	path := "/realms"

	var realms []RealmRepresentation
	_, err := s.client.newRequest(ctx).
		SetResult(&realms).
		Get(path)

	if err != nil {
		return nil, err
	}

	return realms, nil
}

// Create realm with realm, known in Keycloak as import
func (s *RealmService) Create(ctx context.Context, realm *RealmRepresentation) error {
	path := "/realms"
	_, err := s.client.newRequest(ctx).
		SetBody(realm).
		Post(path)

	return err
}

// Delete realm with realm name (not id!)
func (s *RealmService) Delete(ctx context.Context, realm string) error {

	path := "/realms/{realm}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(realm).
		Delete(path)

	return err
}
