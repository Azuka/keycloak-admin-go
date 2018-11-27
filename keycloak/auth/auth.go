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

// Config describes a 2-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs.
type Config struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

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

// Token refreshes the token by using a new client credentials request.
// tokens received this way do not include a refresh token
func (c *tokenSource) Token() (*oauth2.Token, error) {
	v := url.Values{}
	if len(c.conf.Scopes) > 0 {
		v.Set("scope", strings.Join(c.conf.Scopes, " "))
	}
	for k, p := range c.conf.EndpointParams {
		if _, ok := v[k]; ok {
			return nil, fmt.Errorf("keycloak oauth2: cannot overwrite parameter %q", k)
		}
		v[k] = p
	}

	req, err := http.NewRequest("POST", c.conf.TokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if c.conf.ClientID != "" {
		req.SetBasicAuth(c.conf.ClientID, c.conf.ClientSecret)
	}

	r, err := ctxhttp.Do(c.ctx, c.conf.HTTPClient, req)

	if err != nil {
		return nil, err
	}

	// nolint: errcheck
	defer r.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v\nResponse: %s", r.Status, body)
	}

	var tokenRes struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		IDToken      string `json:"id_token"`
		ExpiresIn    int64  `json:"expires_in"` // relative seconds from now
	}

	err = json.Unmarshal(body, &tokenRes)
	if err != nil {
		return nil, err
	}

	tk := &oauth2.Token{
		Expiry:       time.Now().Add(5 * time.Minute),
		AccessToken:  tokenRes.AccessToken,
		RefreshToken: tokenRes.RefreshToken,
		TokenType:    tokenRes.TokenType,
	}
	return tk.WithExtra(body), nil
}
