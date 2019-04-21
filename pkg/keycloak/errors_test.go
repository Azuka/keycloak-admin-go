package keycloak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorTests = []struct {
	Error         Error
	ExpectedValue string
}{
	{
		Error{},
		": 0",
	},
	{
		Error{
			Code:    401,
			Message: "Not authorized",
		},
		"Not authorized: 401",
	},
}

func TestError(t *testing.T) {
	a := assert.New(t)

	for _, tt := range errorTests {
		a.Equal(tt.ExpectedValue, tt.Error.Error())
	}
}
