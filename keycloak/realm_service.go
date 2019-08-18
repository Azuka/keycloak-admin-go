package keycloak

import "context"

// RealmService interacts with all realm resources
type RealmService service

// NewRealmService returns a new user service for working with user resources
// in a realm.
func NewRealmService(c *Client) *RealmService {
	return &RealmService{
		client: c,
	}
}

// Get realm with realm name (not id!)
func (rs *RealmService) Get(ctx context.Context, realm string) (*RealmRepresentation, error) {

	// nolint: goconst
	path := "/realms/{realm}"

	rr := &RealmRepresentation{}

	_, err := rs.client.newRequest(ctx).
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

// Create realm with realm, known in Keycloak as import
func (rs *RealmService) Create(ctx context.Context, realm *RealmRepresentation) error {
	path := "/realms"
	_, err := rs.client.newRequest(ctx).
		SetBody(realm).
		Post(path)

	return err
}

// Clear a realm's user cache
func (rs *RealmService) ClearUserCache(ctx context.Context, realm string) error {
	path := "/realms/{realm}/clear-user-cache"
	_, err := rs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		Post(path)

	return err
}

// Clear a realm's cache
func (rs *RealmService) ClearCache(ctx context.Context, realm string) error {
	path := "/realms/{realm}/clear-realm-cache"
	_, err := rs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		Post(path)

	return err
}

// Delete realm with realm name (not id!)
func (rs *RealmService) Delete(ctx context.Context, realm string) error {

	// nolint: goconst
	path := "/realms/{realm}"

	_, err := rs.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(realm).
		Delete(path)

	return err
}
