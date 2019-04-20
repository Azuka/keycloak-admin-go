package keycloak

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/jarcoal/httpmock.v1"
	"gopkg.in/resty.v1"
)

func ExampleNewUserService() {
	userService := NewUserService(&Client{})
	_, _ = userService.Create(context.TODO(), "my-realm", &UserRepresentation{})
}

type userServiceTests struct {
	userService *UserService
	suite.Suite
}

func (suite *userServiceTests) SetupSuite() {
	c := &Client{
		BaseURL: url.URL{
			Scheme: "https",
			Path:   "",
			Host:   "keycloak.local",
		},
		restClient: resty.New().OnAfterResponse(handleResponse),
	}
	c.Debug()
	suite.userService = NewUserService(c)
}

func (suite *userServiceTests) SetupTest() {
	httpmock.ActivateNonDefault(suite.userService.client.restClient.GetClient())
}

func (suite *userServiceTests) TeardownTest() {
	httpmock.DeactivateAndReset()
}

func TestNewUserService(t *testing.T) {
	a := assert.New(t)
	c := &Client{}
	userService := NewUserService(c)

	a.NotNil(userService)
	a.Equal(c, userService.client)
}

func (suite *userServiceTests) TestUserServiceCreateUser() {
	response := httpmock.NewStringResponse(201, "")
	response.Header.Add("Location", "https://keycloak.local/realms/my-realm/users/my-awesome-id")
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("POST", "https://keycloak.local/realms/my-realm/users", responder)

	id, err := suite.userService.Create(context.TODO(), "my-realm", &UserRepresentation{
		Username: "me",
	})
	suite.NoError(err)
	suite.Equal("my-awesome-id", id)
}

func (suite *userServiceTests) TestUserServiceCreateUserFailure() {
	response := httpmock.NewStringResponse(500, "")
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("POST", "https://keycloak.local/realms/my-realm/users", responder)

	_, err := suite.userService.Create(context.TODO(), "my-realm", &UserRepresentation{
		Username: "me",
	})
	suite.NotNil(err)

	actualError, ok := err.(*Error)

	suite.True(ok)
	suite.NotNil(actualError)
	suite.Equal(500, actualError.Code)
}

func (suite *userServiceTests) TestUserServiceUpdateUser() {
	response := httpmock.NewStringResponse(204, "")
	response.Header.Add("Location", "https://keycloak.local/realms/my-realm/users/my-awesome-id")
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("PUT", "https://keycloak.local/realms/my-realm/users/abc", responder)

	err := suite.userService.Update(context.TODO(), "my-realm", &UserRepresentation{
		Username: "me",
		ID:       "abc",
	})
	suite.NoError(err)
}

func TestUserServiceMethods(t *testing.T) {
	suite.Run(t, &userServiceTests{})
}
