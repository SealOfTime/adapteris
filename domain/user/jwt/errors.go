package jwt

import "fmt"

//ErrSignJwt signals unsuccessful attempt at signing JWT
type ErrSignJwt struct {
	Cause error
}

func (e ErrSignJwt) Error() string {
	return fmt.Sprintf("error while signing jwt: %+v", e.Cause)
}

func (e ErrSignJwt) Unwrap() error {
	return e.Cause
}

//ErrBadJwt signals that JWT is invalid
type ErrBadJwt struct {
	Cause error
}

func (e ErrBadJwt) Error() string {
	return fmt.Sprintf("error while validating jwt: %+v", e.Cause)
}

func (e ErrBadJwt) Unwrap() error {
	return e.Cause
}
