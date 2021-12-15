package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type Authenticator interface {
	//Authenticate returns jwt.Token on successfully verifying the identity of provided credentials or an error in the
	//case of invalid credentials.
	Authenticate(ctx context.Context, credentials interface{}) (*jwt.Token, error)
}

// Authenticate attempts to authenticate provided credentials in every registered Authenticator, obrupting the process
// on any returned error and successfully authenticating user on first returned jwt.Token.
//
// Returns nil, nil in case none of the Authenticators has verified the identity.
func (a *appImpl) Authenticate(ctx context.Context, credentials interface{}) (*jwt.Token, error) {
	for _, p := range a.providers {
		token, err := p.Authenticate(ctx, credentials)
		//If authentication in one of providers ended with error immediately stop authentication process.
		if err != nil {
			return nil, err
		}

		if token != nil {
			return token, nil
		}
	}
	return nil, nil
}
