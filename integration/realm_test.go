package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/pkg/keycloak"
)

func (suite *integrationTester) TestRealmFetch() {
	realm, err := suite.client.Realms().Get(suite.ctx, keycloakAdminRealm)
	suite.NotNil(realm, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(keycloakAdminRealm, realm.ID, suite.version)
}

func (suite *integrationTester) TestRealmDelete() {
	realmID := pseudoRandString()
	realmName := pseudoRandString()

	newRealm := &keycloak.RealmRepresentation{
		ID:    realmID,
		Realm: realmName,
	}

	err := suite.client.Realms().Create(suite.ctx, newRealm)
	suite.NoError(err, suite.version)

	err = suite.client.Realms().Delete(suite.ctx, realmName)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestRealmCreate() {
	realmID := pseudoRandString()
	realmName := pseudoRandString()
	t := func() *bool { b := true; return &b }()
	newRealm := &keycloak.RealmRepresentation{
		ID:                                  realmID,
		Realm:                               realmName,
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

	err := suite.client.Realms().Create(suite.ctx, newRealm)
	suite.NoError(err, suite.version)

	actualRealm, err := suite.client.Realms().Get(suite.ctx, realmName)
	suite.NoError(err, suite.version)
	suite.NotNil(actualRealm, suite.version)
	suite.Equal(actualRealm.ID, newRealm.ID, suite.version)
	suite.Equal(actualRealm.Realm, newRealm.Realm, suite.version)

	suite.Equal(actualRealm.AccessCodeLifespan, newRealm.AccessCodeLifespan, suite.version)
	suite.Equal(actualRealm.AccessCodeLifespanLogin, newRealm.AccessCodeLifespanLogin, suite.version)
	suite.Equal(actualRealm.AccessCodeLifespanUserAction, newRealm.AccessCodeLifespanUserAction, suite.version)
	suite.Equal(actualRealm.AccessTokenLifespan, newRealm.AccessTokenLifespan, suite.version)
	suite.Equal(actualRealm.AccessTokenLifespanForImplicitFlow, newRealm.AccessTokenLifespanForImplicitFlow, suite.version)
	suite.Equal(actualRealm.AccountTheme, newRealm.AccountTheme, suite.version)
	suite.Equal(actualRealm.ActionTokenGeneratedByAdminLifespan, newRealm.ActionTokenGeneratedByAdminLifespan, suite.version)
	suite.Equal(actualRealm.ActionTokenGeneratedByUserLifespan, newRealm.ActionTokenGeneratedByUserLifespan, suite.version)
	suite.Equal(actualRealm.AdminEventsDetailsEnabled, newRealm.AdminEventsDetailsEnabled, suite.version)
	suite.Equal(actualRealm.AdminEventsEnabled, newRealm.AdminEventsEnabled, suite.version)
	suite.Equal(actualRealm.AdminTheme, newRealm.AdminTheme, suite.version)
	suite.Equal(actualRealm.DisplayName, newRealm.DisplayName, suite.version)
	suite.Equal(actualRealm.DisplayNameHTML, newRealm.DisplayNameHTML, suite.version)
}
