package auth

import (
	"context"
	"fmt"
	"os"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()
	providerURL := fmt.Sprintf("https://%s/", os.Getenv("AUTH0_DOMAIN"))
	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		return nil, err
	}
	config := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_REDIRECT_URL"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}
	return &Authenticator{
		Provider: provider,
		Config:   config,
		Ctx:      ctx,
	}, nil
}
