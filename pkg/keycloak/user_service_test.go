package keycloak

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/jarcoal/httpmock.v1"
	"gopkg.in/resty.v1"
)

type userServiceTests struct {
	client *Client
	suite.Suite
	baseURL string
}

func (suite *userServiceTests) SetupSuite() {
	baseUserURL, err := url.Parse(fmt.Sprintf("http://keycloak.local/realms/%s/users", testRelam))
	suite.NoError(err)
	c := &Client{
		Server: url.URL{
			Scheme: "https",
			Path:   "",
			Host:   "keycloak.local",
		},
		restClient: resty.New().OnAfterResponse(handleResponse),
		Realm:      testRelam,
	}
	c.Debug()
	suite.client = c
	suite.baseURL = baseUserURL.Path
}

func (suite *userServiceTests) SetupTest() {
	httpmock.ActivateNonDefault(suite.client.restClient.GetClient())
}

func (suite *userServiceTests) TeardownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *userServiceTests) TestUserServiceCreateUser() {
	response := httpmock.NewStringResponse(201, "")
	response.Header.Add("Location", path.Join(suite.baseURL, "my-awesome-id"))
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("POST", suite.baseURL, responder)

	id, err := suite.client.Users().Create(context.TODO(), &UserRepresentation{
		Username: "me",
	})
	suite.NoError(err)
	suite.Equal("my-awesome-id", id)
}

func (suite *userServiceTests) TestUserServiceCreateUserFailure() {
	response := httpmock.NewStringResponse(500, "")
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("POST", suite.baseURL, responder)

	_, err := suite.client.Users().Create(context.TODO(), &UserRepresentation{
		Username: "me",
	})
	suite.NotNil(err)

	actualError, ok := err.(Error)
	fmt.Println(err.Error())

	suite.True(ok)
	suite.NotNil(actualError)
	suite.Equal(500, actualError.Code)
}

func (suite *userServiceTests) TestUserServiceUpdateUser() {
	response := httpmock.NewStringResponse(204, "")
	response.Header.Add("Location", path.Join(suite.baseURL, "my-awesome-id"))
	responder := httpmock.ResponderFromResponse(response)

	fmt.Println(path.Join(suite.baseURL, "abc"))
	httpmock.RegisterResponder("PUT", path.Join(suite.baseURL, "abc"), responder)

	err := suite.client.Users().Update(context.TODO(), &UserRepresentation{
		Username: "me",
		ID:       "abc",
	})
	suite.NoError(err)
}

func TestUserServiceMethods(t *testing.T) {
	suite.Run(t, &userServiceTests{})
}
