package integration_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/Azuka/keycloak-admin-go/keycloak"
	"github.com/Azuka/keycloak-admin-go/keycloak/auth"
	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/suite"
)

const keycloakEndpoint = "http://127.0.0.1:9090/auth/"
const keycloakAdmin = "keycloak-admin"
const keycloakPassword = "changeme"
const keycloakAdminRealm = "master"
const keycloakAdminClientID = "admin-cli"

type integrationTester struct {
	ready chan struct{}
	suite.Suite
	client *keycloak.Client
	ctx    context.Context
}

func getHTTPClient(ctx context.Context) *http.Client {

	config := auth.Config{
		ClientID: keycloakAdminClientID,
		TokenURL: keycloakEndpoint + "realms/" + keycloakAdminRealm + "/protocol/openid-connect/token",
		EndpointParams: url.Values{
			"username":   {keycloakAdmin},
			"password":   {keycloakPassword},
			"grant_type": {"password"},
			"client_id":  {keycloakAdminClientID},
		},
	}

	http.DefaultClient.Timeout = time.Second * 5

	return config.Client(ctx)
}

func (suite *integrationTester) SetupSuite() {
	suite.ready = make(chan struct{})
	suite.ctx = context.Background()

	connect := func() error {
		_, err := http.Get(keycloakEndpoint)

		if err == nil {
			close(suite.ready)
			return nil
		}

		fmt.Println("Waiting to connect to keycloak: ", err)

		return err
	}

	// Setup test client
	u, _ := url.Parse(keycloakEndpoint + "admin")
	suite.client = keycloak.NewClient(*u, getHTTPClient(suite.ctx))
	suite.client.Debug()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	go func() {
		err := backoff.Retry(connect, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
		if err != nil {
			fmt.Println("error connecting: ", err)
		}
	}()

	<-suite.ready
}

func TestKeycloakAdminIntegration(t *testing.T) {
	suite.Run(t, &integrationTester{})
}
