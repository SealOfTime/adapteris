package vk

import "fmt"

//ErrProviderNotFound is error that signals, that Provider is not registered in this auth.App
type ErrProviderNotFound struct{ Provider string }

func (e ErrProviderNotFound) Error() string {
	return fmt.Sprintf("provider %s not found", e.Provider)
}

//ErrExchangeFailed signals that certain part of oAuth2 Authorization Code flow has failed.
type ErrExchangeFailed struct{ Cause error }

func (e ErrExchangeFailed) Error() string {
	return fmt.Sprintf("error while exchanging access code: %+v", e.Cause)
}

func (e ErrExchangeFailed) Unwrap() error {
	return e.Cause
}

//ErrOA2InvalidToken signals that something about the OA2 returned to oAuth2Authenticator after exchange is wrong.
type ErrOA2InvalidToken struct{ Cause error }

func (e ErrOA2InvalidToken) Error() string {
	return fmt.Sprintf("error while processing OAuth2 token: %+v", e.Cause)
}

func (e ErrOA2InvalidToken) Unwrap() error {
	return e.Cause
}

//ErrOA2ExtraMissing signals that the expected Extra key-value pair in the oAuth2 token is missing.
type ErrOA2ExtraMissing struct {
	Extra string
}

func (e ErrOA2ExtraMissing) Error() string {
	return e.Extra
}

type ErrVkFetchUserDetails struct {
	Cause error
}

func (e ErrVkFetchUserDetails) Error() string {
	return fmt.Sprintf("error fetching vk user's details: %+v", e.Cause)
}

func (e ErrVkFetchUserDetails) Unwrap() error {
	return e.Cause
}
