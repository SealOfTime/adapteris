package auth

import (
	"context"
	"fmt"

	domain "github.com/sealoftime/adapteris"
)

type Authenticator interface {
	Authenticate(context.Context, interface{}) (*domain.User, error)
}

type App struct {
	loginUrlFunc func(redirectUrl string) string
	methods      map[string]Authenticator
}

func NewApp(
	loginUrlFunc func(redirectUrl string) string,
	methods map[string]Authenticator,
) App {
	return App{loginUrlFunc: loginUrlFunc, methods: methods}
}

func (a *App) ConcatLoginUrl(redirect string) string {
	return a.loginUrlFunc(redirect)
}

func (a *App) Authenticate(ctx context.Context, creds interface{}) (*domain.User, error) {
	for _, method := range a.methods {
		u, err := method.Authenticate(ctx, creds)
		if err != nil {
			return nil, fmt.Errorf("authentication error: %w", err)
		}
		if u != nil {
			return u, nil
		}
	}

	return nil, nil
}
