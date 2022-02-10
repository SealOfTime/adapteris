package jwt

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	domain "github.com/sealoftime/adapteris"
)

type Authenticator struct {
	secret      []byte
	userStorage domain.UserStorage
}

type Creds = string

func (a Authenticator) Authenticate(ctx context.Context, rawCreds interface{}) (*domain.User, error) {
	creds, ok := rawCreds.(Creds)
	if !ok {
		return nil, nil
	}

	var claims jwt.StandardClaims
	t, err := jwt.ParseWithClaims(creds, &claims, a.getSigningKeyForToken)
	switch {
	case err != nil:
		return nil, ErrBadJwt{err}
	case !t.Valid:
		return nil, ErrBadJwt{fmt.Errorf("Token invalid")}
	}

	uid, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, ErrBadJwt{fmt.Errorf("Subject is not a user id")}
	}

	u, err := a.userStorage.FindById(ctx, uid)
	if err != nil {
		return nil, ErrBadJwt{err}
	}

	return u, nil
}

func (a Authenticator) getSigningKeyForToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return a.secret, nil
}
