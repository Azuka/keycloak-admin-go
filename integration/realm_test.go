package integration_test

import (
	"github.com/Azuka/keycloak-admin-go/keycloak"
	"github.com/satori/go.uuid"
)

func (suite *integrationTester) TestRealmFetch() {
	realm, err := suite.client.Realm.Get(suite.ctx, keycloakAdminRealm)
	suite.NotNil(realm)
	suite.NoError(err)
	suite.Equal(keycloakAdminRealm, realm.ID)
}

func (suite *integrationTester) TestRealmDelete() {
	realmID, _ := uuid.NewV4()
	realmName, _ := uuid.NewV4()

	newRealm := &keycloak.RealmRepresentation{
		ID:    realmID.String(),
		Realm: realmName.String(),
	}

	err := suite.client.Realm.Create(suite.ctx, newRealm)
	suite.NoError(err)

	err = suite.client.Realm.Delete(suite.ctx, realmName.String())
	suite.NoError(err)
}

func (suite *integrationTester) TestRealmCreate() {
	realmID, _ := uuid.NewV4()
	realmName, _ := uuid.NewV4()
	t := func() *bool { b := true; return &b }()
	newRealm := &keycloak.RealmRepresentation{
		ID:                                  realmID.String(),
		Realm:                               realmName.String(),
		AccessCodeLifespan:                  1,
		AccessCodeLifespanLogin:             2,
		AccessCodeLifespanUserAction:        3,
		AccessTokenLifespan:                 4,
		AccessTokenLifespanForImplicitFlow:  5,
		AccountTheme:                        "base",
		ActionTokenGeneratedByAdminLifespan: 6,
		ActionTokenGeneratedByUserLifespan:  7,
		AdminEventsDetailsEnabled:           t,
		AdminEventsEnabled:                  t,
		AdminTheme:                          "base",
		DisplayName:                         "realmDisplayName",
		DisplayNameHTML:                     "realmDisplayNameHTML",
	}

	err := suite.client.Realm.Create(suite.ctx, newRealm)
	suite.NoError(err)

	actualRealm, err := suite.client.Realm.Get(suite.ctx, realmName.String())
	suite.NoError(err)
	suite.NotNil(actualRealm)
	suite.Equal(actualRealm.ID, newRealm.ID)
	suite.Equal(actualRealm.Realm, newRealm.Realm)

	suite.Equal(actualRealm.AccessCodeLifespan, newRealm.AccessCodeLifespan)
	suite.Equal(actualRealm.AccessCodeLifespanLogin, newRealm.AccessCodeLifespanLogin)
	suite.Equal(actualRealm.AccessCodeLifespanUserAction, newRealm.AccessCodeLifespanUserAction)
	suite.Equal(actualRealm.AccessTokenLifespan, newRealm.AccessTokenLifespan)
	suite.Equal(actualRealm.AccessTokenLifespanForImplicitFlow, newRealm.AccessTokenLifespanForImplicitFlow)
	suite.Equal(actualRealm.AccountTheme, newRealm.AccountTheme)
	suite.Equal(actualRealm.ActionTokenGeneratedByAdminLifespan, newRealm.ActionTokenGeneratedByAdminLifespan)
	suite.Equal(actualRealm.ActionTokenGeneratedByUserLifespan, newRealm.ActionTokenGeneratedByUserLifespan)
	suite.Equal(actualRealm.AdminEventsDetailsEnabled, newRealm.AdminEventsDetailsEnabled)
	suite.Equal(actualRealm.AdminEventsEnabled, newRealm.AdminEventsEnabled)
	suite.Equal(actualRealm.AdminTheme, newRealm.AdminTheme)
	suite.Equal(actualRealm.DisplayName, newRealm.DisplayName)
	suite.Equal(actualRealm.DisplayNameHTML, newRealm.DisplayNameHTML)
}
