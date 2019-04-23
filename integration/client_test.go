package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/pkg/keycloak"
)

func (suite *integrationTester) TestClientCreate() {

	client := &keycloak.ClientRepresentation{
		ClientID: pseudoRandString(),
	}

	id, err := suite.client.Clients().Create(client)

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestClientFetch() {
	clientName := "admin-cli"
	clients, err := suite.client.Clients().Find(map[string]string{
		"clientId": clientName,
	})
	suite.NotNil(clients, suite.version)
	suite.NoError(err, suite.version)
	suite.Len(clients, 1, suite.version)

	client, err := suite.client.Clients().Get(clients[0].ID)
	suite.NotNil(clients, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(clientName, client.ClientID, suite.version)
}
