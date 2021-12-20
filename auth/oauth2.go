package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"time"
)

type OAuth2AccessCodeCredentials struct {
	AccessCode   string
	ProviderName string
}

type ExternalIdFromOA2Func func(ctx context.Context, oa2t *oauth2.Token) (string, error)

type oAuth2Authenticator struct {
	//Name is a unique string identifying this Authenticator from others
	Name                   string
	ExtractExternalId      ExternalIdFromOA2Func
	ExternalAccountStorage ExternalAccountStorage
	Cfg                    ProviderConfig
}

var _ Authenticator = (*oAuth2Authenticator)(nil)

//NewOAuth2Authenticator returns new Authenticator, that validates the identity of given credentials by exchanging
//AccessCode from OAuth2AccessCodeCredentials for a token of external service, after which it calls extIdFunc to
//extract the ExternalAccount's id and then generate a jwt.Token for AdapterIS with User's id discovered from the ExternalAccountStorage.
func NewOAuth2Authenticator(
	name string,
	cfg *ProviderConfig,
	extIdFunc ExternalIdFromOA2Func,
	storage ExternalAccountStorage,
) *oAuth2Authenticator {
	return &oAuth2Authenticator{
		Cfg:                    *cfg,
		Name:                   name,
		ExtractExternalId:      extIdFunc,
		ExternalAccountStorage: storage,
	}
}

//Authenticate implements Authenticator.Authenticate function
//
//• ErrExchangeFailed propagates error from oauth2.Config's Exchange
//
//• ErrOA2InvalidToken propagates error from oAuth2Authenticator.ExtractExternalId
//
func (a *oAuth2Authenticator) Authenticate(ctx context.Context, credentials interface{}) (string, error) {
	creds, ok := credentials.(OAuth2AccessCodeCredentials)
	if !ok || creds.ProviderName != a.Name {
		//If not the OAuth2 credentials authenticating or if the requested authentication provider is not the same
		//as this one, skip.
		return "", nil
	}

	oa2t, err := a.Cfg.OAuth2.Exchange(ctx, creds.AccessCode)
	if err != nil {
		return "", ErrExchangeFailed{Cause: err}
	}

	eid, err := a.ExtractExternalId(ctx, oa2t)
	if err != nil {
		return "", ErrOA2InvalidToken{Cause: err}
	}

	ea, err := a.ExternalAccountStorage.FindByExternalId(ctx, eid)
	if err != nil {
		return "", err
	}

	now := time.Now()
	rawT := jwt.NewWithClaims(a.Cfg.Jwt.signingMethod, jwt.RegisteredClaims{
		Subject:   fmt.Sprintf(`{"uid": "%d"}`, ea.UserId),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * 20)),
		IssuedAt:  jwt.NewNumericDate(now),
	})
	t, err := rawT.SignedString(a.Cfg.Jwt.key)
	if err != nil {
		return "", ErrSignJwt{Cause: err}
	}

	return t, nil
}
