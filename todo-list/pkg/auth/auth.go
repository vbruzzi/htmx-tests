package auth

import (
	"context"
	"vbruzzi/todo-list/pkg/config"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func New(ctx context.Context, config config.OidcConfig) {
	provider, err := oidc.NewProvider(ctx, config.Authority)
	if err != nil {
		// handle error
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectUrl,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
}
