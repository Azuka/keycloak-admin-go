package keycloak

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func ExampleNewUserService() {
	userService := NewUserService(&Client{})
	userService.Create(context.TODO(), "my-realm", &UserRepresentation{})
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
		restClient: resty.New(),
	}
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
	response.Header.Add("Location", "https://keycloak.local/my-realm/users/my-awesome-id")
	responder := httpmock.ResponderFromResponse(response)

	httpmock.RegisterResponder("POST", "https://keycloak.local/my-realm/users", responder)

	id, err := suite.userService.Create(context.TODO(), "my-realm", &UserRepresentation{
		Username: "me",
	})
	suite.NoError(err)
	suite.Equal("my-awesome-id", id)
}

func TestUserServiceMethods(t *testing.T) {
	suite.Run(t, &userServiceTests{})
}
