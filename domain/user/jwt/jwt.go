package jwt

import (
	"context"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/sealoftime/adapteris/user"
)

type Service struct {
	method   jwt.SigningMethod
	accounts user.Repository
	secret   []byte
}

var (
	defaultSigningMethod = jwt.SigningMethodHS512
)

func NewService(
	accountStorage user.Repository,
	secret []byte,
) Service {
	return Service{
		method:   defaultSigningMethod,
		accounts: accountStorage,
		secret:   secret,
	}
}

func (s *Service) SignAccessToken(
	ctx context.Context,
	user user.Account,
) (tokenString string, err error) {
	token := jwt.NewWithClaims(s.method, jwt.StandardClaims{
		Subject: strconv.FormatInt(user.Id, 10),
	})

	return token.SignedString(s.secret)
}
