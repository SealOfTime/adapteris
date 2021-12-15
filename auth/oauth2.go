package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

type OAuth2AccessCodeCredentials struct {
	AccessCode   string
	ProviderName string
}

type JWTFromOA2TokenFunc func(ctx context.Context, oa2t *oauth2.Token) (*jwt.Token, error)

type oAuth2Authenticator struct {
	oauth2.Config
	Name string
	//tokenConverter must be implemented separately or otherwise
	Convert JWTFromOA2TokenFunc
}

var _ Authenticator = (*oAuth2Authenticator)(nil)

//NewOAuth2Authenticator returns new Authenticator, that validates the identity of given credentials by exchanging
//AccessCode from OAuth2AccessCodeCredentials for a token of external service, after which it calls convertFunc to
//generate a jwt.Token for AdapterIS.
func NewOAuth2Authenticator(name string, cfg oauth2.Config, convertFunc JWTFromOA2TokenFunc) *oAuth2Authenticator {
	return &oAuth2Authenticator{
		Config: cfg,
		Name: name,
		Convert: convertFunc,
	}
}

//Authenticate implements Authenticator.Authenticate function
func (a *oAuth2Authenticator) Authenticate(ctx context.Context, credentials interface{}) (*jwt.Token, error) {
	creds, ok := credentials.(OAuth2AccessCodeCredentials)
	if !ok || creds.ProviderName != a.Name {
		//If not the OAuth2 credentials authenticating or if the requested authentication provider is not the same
		//as this one, skip.
		return nil, nil
	}

	t, err := a.Exchange(ctx, creds.AccessCode)
	if err != nil {
		return nil, ErrExchangeFailed{Cause: err}
	}

	return a.Convert(ctx, t)
}