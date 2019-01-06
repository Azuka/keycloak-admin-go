package integration_test

import (
	"github.com/Azuka/keycloak-admin-go/keycloak"
)

func (suite *integrationTester) TestUserFetch() {
	users, err := suite.client.Users.Find(suite.ctx, keycloakAdminRealm, map[string]string{
		"username": keycloakAdmin,
	})
	suite.NotNil(users, suite.version)
	suite.NoError(err, suite.version)
	suite.Len(users, 1, suite.version)
	suite.Equal(keycloakAdmin, users[0].Username, suite.version)
	suite.True(*users[0].Enabled, suite.version)

	user := users[0]
	t := true
	user.EmailVerified = &t

	err = suite.client.Users.Update(suite.ctx, keycloakAdminRealm, &user)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestUserCreate() {

	user := &keycloak.UserRepresentation{
		Username: pseudoRandString(),
		Email:    pseudoRandString() + "@example.com",
	}

	id, err := suite.client.Users.Create(suite.ctx, keycloakAdminRealm, user)

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
}
