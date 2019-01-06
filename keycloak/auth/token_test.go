package auth

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"testing"
)

func TestToken_Oauth2Token(t *testing.T) {
	a := assert.New(t)

	tkn := &Token{}
	otkn := tkn.Oauth2Token()

	a.Equal(tkn, Extract(otkn))
}

func TestToken_Oauth2Token_Empty(t *testing.T) {
	a := assert.New(t)

	tkn := &oauth2.Token{}
	a.Nil(Extract(tkn))
}
