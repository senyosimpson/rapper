package spotify

import (
	"context"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

// ClientCredentialsAuth is a struct that contains a client credentials config
type ClientCredentialsAuth struct {
	config *clientcredentials.Config
}

// NewClientCredentialsAuth creates a new client credentials authenticator.
func NewClientCredentialsAuth() ClientCredentialsAuth {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		TokenURL:     TokenURL,
	}
	return ClientCredentialsAuth{config: config}
}

// NewClient creates a new client from the credentials authenticator
func (auth ClientCredentialsAuth) NewClient(ctx context.Context) Client {
	client := auth.config.Client(ctx)
	return Client{
		client:  client,
		baseURL: BaseURL,
	}
}
