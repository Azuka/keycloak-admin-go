package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/pkg/keycloak"
)

func (suite *integrationTester) TestRolesCreate() {

	role := &keycloak.RoleRepresentation{
		Name: pseudoRandString(),
	}

	id, err := suite.client.Roles().Create("master", role)

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
}
