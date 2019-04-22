package auth

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2"
)

// Config is the configuration for an OAuth client
type OAuth struct {
	Server      url.URL
	ID          string
	Secret      string
	RedirectURL string
	Realm       string
}

// PasswordCredentialsClient gets an http client from username and password
func (c *OAuth) PasswordCredentialsClient(username, password string) (*http.Client, error) {
	config, err := c.getConfig()
	if err != nil {
		return nil, err
	}
	ctx := c.getContext()
	token, err := config.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Authentication failed")
		return nil, err
	}
	return config.Client(ctx, token), nil
}

// getContext builds a new context for use in an oauth exchange
func (c *OAuth) getContext() context.Context {
	client := &http.Client{}
	http.DefaultClient.Timeout = time.Second * 5
	return context.WithValue(context.Background(), oauth2.HTTPClient, client)
}

// getConfig builds the oauth2 config for an OAuth client
func (c *OAuth) getConfig() (*oauth2.Config, error) {
	tokenURL := c.Server
	tokenURL.Path = path.Join(tokenURL.Path, "/auth/realms", c.Realm, "/protocol/openid-connect/token")
	authURL := c.Server
	authURL.Path = path.Join(authURL.Path, "/auth/realms", c.Realm, "/protocol/openid-connect/authorize")
	config := &oauth2.Config{
		ClientID:     c.ID,
		ClientSecret: c.Secret,
		RedirectURL:  c.RedirectURL,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL.String(),
			AuthURL:  authURL.String(),
		},
	}
	return config, nil
}
