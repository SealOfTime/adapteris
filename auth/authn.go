package auth

import (
	"context"
)

type Authenticator interface {
	//Authenticate returns jwt.Token on successfully verifying the identity of provided credentials or an error in the
	//case of invalid credentials.
	Authenticate(ctx context.Context, credentials interface{}) (string, error)
}

// Authenticate attempts to authenticate provided credentials in every registered Authenticator, obrupting the process
// on any returned error and successfully authenticating user on first returned jwt.Token.
//
// Returns "", nil in case none of the Authenticators has verified the identity.
func (a *appImpl) Authenticate(ctx context.Context, credentials interface{}) (string, error) {
	for _, p := range a.providers {
		token, err := p.Authenticate(ctx, credentials)
		//If authentication in one of providers ended with error immediately stop authentication process.
		if err != nil {
			return "", err
		}

		if token != "" {
			return token, nil
		}
	}
	return "", nil
}
