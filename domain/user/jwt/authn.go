package jwt

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/sealoftime/adapteris/domain/user"
)

func (s *Service) Authenticate(ctx context.Context, token string) (*user.Account, error) {
	var claims jwt.StandardClaims
	t, err := jwt.ParseWithClaims(token, &claims, s.getSigningKeyForToken)
	switch {
	case err != nil:
		return nil, ErrBadJwt{err}
	case !t.Valid:
		return nil, ErrBadJwt{fmt.Errorf("token invalid")}
	}

	uid, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, ErrBadJwt{fmt.Errorf("subject is not a user id")}
	}

	u, err := s.accounts.FindById(ctx, uid)
	if err != nil {
		return nil, ErrBadJwt{err}
	}

	return u, nil
}

func (s *Service) getSigningKeyForToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return s.secret, nil
}
