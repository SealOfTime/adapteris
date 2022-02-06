package vk

import (
	"context"

	domain "github.com/sealoftime/adapteris"
)

type OAuth2AccessCodeCredentials struct {
	AccessCode   string
	ProviderName string
}

type Authenticator interface {
	Authenticate(context.Context, OAuth2AccessCodeCredentials) (*domain.User, error)
	LoginUrl(afterLoginUrl string) string
}
