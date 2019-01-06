package auth

import (
	"golang.org/x/oauth2"
	"time"
)

type TokenSource interface {
	oauth2.TokenSource

	// KeycloakToken returns a keycloak token
	KeycloakToken() (*Token, error)
}

type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// TokenType is the type of token.
	// The Type method returns either this or "Bearer", the default.
	TokenType string `json:"token_type,omitempty"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refresh_token,omitempty"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry time.Time `json:"expiry,omitempty"`

	// ExpiresIn is the time this token is valid for, per Keycloak
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// RefreshExpiresIn is the time the refresh token expires
	RefreshExpiresIn int64 `json:"refresh_expires_in,omitempty"`

	// NotBeforePolicy is likely the Keycloak clock skew
	NotBeforePolicy int64 `json:"not_before_policy,,omitempty"`

	// SessionState means something in keycloak
	SessionState string `json:"session_state,omitempty"`

	// Scope is the token scope
	Scope string `json:"scope,omitempty"`
}

// Oauth2Token returns an oauth2 token with the underlying original keycloak token
func (t *Token) Oauth2Token() *oauth2.Token {

	tkn := &oauth2.Token{
		AccessToken: t.AccessToken,
		TokenType:   t.TokenType,
		RefreshToken:   t.RefreshToken,
		Expiry: t.Expiry,
	}

	return tkn.WithExtra(t)
}
