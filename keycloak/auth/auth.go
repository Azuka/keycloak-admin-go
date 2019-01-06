// Package auth is copied almost verbatim from golang.org/x/oauth2/clientcredentials
//
// This is because the package above doesn't allow overwriting the grant_type key
// TODO: Clean up and implement/reuse a true keycloak auth
package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"encoding/json"
	"io"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/oauth2"
)

const (
	// ClientCredentialsGrant applies to client credentials
	ClientCredentialsGrant = "client_credentials"

	// PasswordGrant is for the password grant
	PasswordGrant = "password"
)

// skew for token expiry
const expirationSkew = 5

// Config describes a 2-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
type Config struct {
	// ClientID is the application's ID. This should be set for both
	// password and client credentials grants
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// Username is the username (if using the password grant).
	Username string

	// Password is user's password (if using the password grant).
	Password string

	// GrantType is the auth grant type
	GrantType string

	// TokenURL is the resource server's token endpoint
	// URL. This is a constant specific to each server.
	TokenURL string

	// Scope specifies optional requested permissions.
	Scopes []string

	// EndpointParams specifies additional parameters for requests to the token endpoint.
	EndpointParams url.Values

	HTTPClient *http.Client
}

// Token uses client credentials to retrieve a token.
// The HTTP client to use is derived from the context.
// If nil, http.DefaultClient is used.
func (c *Config) Token(ctx context.Context) (*oauth2.Token, error) {
	return c.TokenSource(ctx).Token()
}

// Client returns an HTTP client using the provided token.
// The token will auto-refresh as necessary. The underlying
// HTTP transport will be obtained using the provided context.
// The returned client and its Transport should not be modified.
func (c *Config) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx))
}

// TokenSource returns a TokenSource that returns t until t expires,
// automatically refreshing it as necessary using the provided context and the
// client ID and client secret.
//
// Most users will use Config.Client instead.
func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
	source := &tokenSource{
		ctx:  ctx,
		conf: c,
	}
	return oauth2.ReuseTokenSource(nil, source)
}

type tokenSource struct {
	ctx  context.Context
	conf *Config
}

// KeycloakToken refreshes the token by using a new request.
// tokens received this way do not include a refresh token
func (c *tokenSource) KeycloakToken() (*Token, error) {
	v := url.Values{}

	// Set scopes
	if len(c.conf.Scopes) > 0 {
		v.Set("scope", strings.Join(c.conf.Scopes, " "))
	}

	// Set client_id and client_secret
	if c.conf.ClientID != "" {
		v.Set("client_id", c.conf.ClientID)
	}
	if c.conf.ClientSecret != "" {
		v.Set("client_secret", c.conf.ClientSecret)
	}

	// Set grant type
	if c.conf.GrantType != "" {
		v.Set("grant_type", c.conf.GrantType)
	}

	// Set username and password
	if c.conf.Username != "" {
		v.Set("username", c.conf.Username)
	}
	if c.conf.Password != "" {
		v.Set("password", c.conf.Password)
	}

	for k, p := range c.conf.EndpointParams {
		if _, ok := v[k]; ok {
			return nil, fmt.Errorf("keycloak oauth2: cannot overwrite parameter %q", k)
		}
		v[k] = p
	}

	req, err := http.NewRequest("POST", c.conf.TokenURL, strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	r, err := ctxhttp.Do(c.ctx, c.conf.HTTPClient, req)

	if err != nil {
		return nil, err
	}

	if r.Body == nil {
		return nil, fmt.Errorf("oauth2: empty keycloak auth response")
	}

	// nolint: errcheck
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch keycloak token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, fmt.Errorf("oauth2: cannot fetch keycloak token: %v\nResponse: %s", r.Status, body)
	}

	tk := &Token{}

	err = json.Unmarshal(body, tk)

	if err != nil {
		return nil, err
	}

	tk.Expiry = time.Now().Add(time.Second * time.Duration(tk.ExpiresIn-expirationSkew))

	return tk, nil
}

// Token returns the oauth2.Token representation of the keycloak token
func (c *tokenSource) Token() (*oauth2.Token, error) {

	tkn, err := c.KeycloakToken()

	if err != nil {
		return nil, err
	}

	return tkn.Oauth2Token(), nil
}
