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
