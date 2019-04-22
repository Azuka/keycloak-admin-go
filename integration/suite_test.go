package integration_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/Azuka/keycloak-admin-go/pkg/auth"
	"github.com/Azuka/keycloak-admin-go/pkg/keycloak"
	"github.com/cenkalti/backoff"
	"github.com/stretchr/testify/suite"
)

const keycloakAdmin = "keycloak-admin"
const keycloakPassword = "changeme"
const keycloakAdminRealm = "master"
const keycloakAdminClientID = "admin-cli"

var keyCloakEndpoints = map[string]string{
	"5.0.0": "http://127.0.0.1:9090",
	"4.8.0": "http://127.0.0.1:9098",
}

type integrationTester struct {
	ready chan struct{}
	suite.Suite
	client   *keycloak.Client
	ctx      context.Context
	version  string
	endpoint string
}

func (suite *integrationTester) httpClient() *http.Client {
	server, err := url.Parse(suite.endpoint)
	if err != nil {
		close(suite.ready)
		return nil
	}
	oauth := &auth.OAuth{
		Server: *server,
		ID:     keycloakAdminClientID,
		Realm:  keycloakAdminRealm,
	}
	client, err := oauth.PasswordCredentialsClient(keycloakAdmin, keycloakPassword)
	if err != nil {
		close(suite.ready)
		return nil
	}

	return client
}

func (suite *integrationTester) SetupSuite() {
	suite.ready = make(chan struct{})
	suite.ctx = context.Background()

	connect := func() error {
		_, err := http.Get(suite.endpoint)

		if err == nil {
			close(suite.ready)
			return nil
		}

		fmt.Println("Waiting to connect to keycloak: ", err)

		return err
	}

	// Setup test client
	u, _ := url.Parse(suite.endpoint)
	suite.client = keycloak.NewClient(*u, suite.httpClient(), keycloakAdminRealm)
	suite.client.Debug()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	go func() {
		err := backoff.Retry(connect, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
		if err != nil {
			panic(fmt.Errorf("error connecting: %+v", err))
		}
	}()

	<-suite.ready
}

func TestKeycloakAdminIntegration(t *testing.T) {
	for version, endpoint := range keyCloakEndpoints {
		suite.Run(t, &integrationTester{
			version:  version,
			endpoint: endpoint,
		})
	}
}
