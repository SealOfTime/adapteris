package auth

import (
	"context"

	domain "github.com/sealoftime/adapteris"
)

type AuthService interface {
	Authenticate(ctx context.Context, credentials interface{}) (*domain.User, error)
}
