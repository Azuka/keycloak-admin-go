package integration_test

import (
	"github.com/Azuka/keycloak-admin-go/keycloak"
)

func (suite *integrationTester) TestUserFetch() {
	users, err := suite.client.Users.Find(suite.ctx, keycloakAdminRealm, map[string]string{
		"username": keycloakAdmin,
	})
	suite.NotNil(users)
	suite.NoError(err)
	suite.Len(users, 1)
	suite.Equal(keycloakAdmin, users[0].Username)
	suite.True(*users[0].Enabled)

	user := users[0]
	t := true
	user.EmailVerified = &t

	err = suite.client.Users.Update(suite.ctx, keycloakAdminRealm, &user)
	suite.NoError(err)
}

func (suite *integrationTester) TestUserCreate() {

	randString, _ := uuid.NewV4()

	user := &keycloak.UserRepresentation{
		Username: pseudoRandString(),
		Email:    pseudoRandString() + "@example.com",
	}

	id, err := suite.client.Users.Create(suite.ctx, keycloakAdminRealm, user)

	suite.NotEmpty(id)
	suite.NoError(err)
}
