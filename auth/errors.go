package auth

import "fmt"

//ErrProviderNotFound is error that signals, that Provider is not registered in this auth.App
type ErrProviderNotFound struct{ Provider string }

func (e ErrProviderNotFound) Error() string {
	return fmt.Sprintf("provider %s not found", e.Provider)
}

//ErrExchangeFailed signals that certain part of oAuth2 Authorization Code flow has failed.
type ErrExchangeFailed struct { Cause error }

func (e ErrExchangeFailed) Error() string {
	return fmt.Sprintf("error while exchanging access code: %+v", e.Cause)
}

func (e ErrExchangeFailed) Unwrap() error {
	return e.Cause
}

