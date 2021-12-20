// Package auth provides basic oAuth2 and local JWT token authentication capabilities.
package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type App interface {
	Authenticate(ctx context.Context, credentials interface{}) (string, error)
	Refresh(ctx context.Context) (jwt.Token, error)
}

type WithFunc func(*App)

func New(cfg *Config, storage ExternalAccountStorage) App {
	app := &appImpl{
		providers: setupProviders(cfg, storage),
	}
	return app
}

type appImpl struct {
	providers []Authenticator
}

func (a *appImpl) Refresh(ctx context.Context) (jwt.Token, error) {
	panic("implement me")
}
