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

const keycloakAdmin = "keycloak-admin"
const keycloakPassword = "changeme"
const keycloakAdminRealm = "master"
const keycloakAdminClientID = "admin-cli"

var keyCloakEndpoints = map[string]string{
	"4.0.0": "http://127.0.0.1:9090/auth/",
	"4.8.0": "http://127.0.0.1:9098/auth/",
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
	config := auth.Config{
		ClientID:  keycloakAdminClientID,
		Username:  keycloakAdmin,
		Password:  keycloakPassword,
		GrantType: auth.PasswordGrant,
		TokenURL:  suite.endpoint + "realms/" + keycloakAdminRealm + "/protocol/openid-connect/token",
	}

	http.DefaultClient.Timeout = time.Second * 5
	return config.Client(suite.ctx)
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
	u, _ := url.Parse(suite.endpoint + "admin")
	suite.client = keycloak.NewClient(*u, suite.httpClient())
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
