package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	Authenticate() (jwt.Token, error)
	Refresh() (jwt.Token, error)
}
