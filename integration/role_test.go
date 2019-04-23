package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/pkg/keycloak"
)

func (suite *integrationTester) TestRolesCreate() {

	role := &keycloak.RoleRepresentation{
		Name: pseudoRandString(),
	}

	id, err := suite.client.Roles().Create(suite.ctx, role)

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestGetComposites() {

	role := &keycloak.RoleRepresentation{
		Name: "admin",
	}

	composites, err := suite.client.Roles().GetComposites(suite.ctx, role)

	suite.NotEmpty(composites, suite.version)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestModifyComposites() {

	// Create a new realm
	role := &keycloak.RoleRepresentation{
		Name: pseudoRandString(),
	}
	_, err := suite.client.Roles().Create(suite.ctx, role)
	suite.NoError(err, suite.version)

	roleToAdd, err := suite.client.Roles().Get(suite.ctx, "admin")
	suite.NoError(err, suite.version)
	composites := []keycloak.RoleRepresentation{*roleToAdd}

	err = suite.client.Roles().AddComposite(suite.ctx, role, composites)
	suite.NoError(err, suite.version)

	err = suite.client.Roles().RemoveComposite(suite.ctx, role, composites)
	suite.NoError(err, suite.version)

	composites, err = suite.client.Roles().GetComposites(suite.ctx, role)
	suite.Empty(composites, suite.version)
	suite.NoError(err, suite.version)

	err = suite.client.Roles().Delete(suite.ctx, role)
	suite.NoError(err, suite.version)
}
