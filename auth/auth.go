// Package auth provides basic oAuth2 and local JWT token authentication capabilities.
package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type App interface {
	Authenticate(ctx context.Context, credentials interface{}) (*jwt.Token, error)
	Refresh(ctx context.Context) (jwt.Token, error)
}

type AppOption func(*App)

func New() (App) {
	return &appImpl{
		providers: make([]Authenticator, 3),
	}
}

type appImpl struct {
	services struct {

	}
	providers []Authenticator
}


func (a *appImpl) Refresh(ctx context.Context) (jwt.Token, error) {
	panic("implement me")
}