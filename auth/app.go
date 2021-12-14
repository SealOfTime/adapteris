// Package auth provides basic oAuth2 and local JWT token authentication capabilities.
package auth

import "github.com/golang-jwt/jwt/v4"

type App interface {
	Authenticate() (jwt.Token, error)
	Refresh() (jwt.Token, error)
}

func New() (App) {
	return &appImpl{

	}
}
type appImpl struct {
	services struct {

	}
}

func (a *appImpl) Authenticate() (jwt.Token, error) {
	panic("implement me") //todo: reconsider the contract for this function.
}

func (a *appImpl) Refresh() (jwt.Token, error) {
	panic("implement me")
}