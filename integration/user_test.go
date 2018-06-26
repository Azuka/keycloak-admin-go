package integration_test

func (suite *integrationTester) TestUserFetch() {
	users, err := suite.client.Users.Find(suite.ctx, keycloakAdminRealm, map[string]string{
		"username": keycloakAdmin,
	})
	suite.NotNil(users)
	suite.NoError(err)
	suite.Len(users, 1)
	suite.Equal(keycloakAdmin, users[0].Username)
	suite.True(users[0].Enabled)
}
