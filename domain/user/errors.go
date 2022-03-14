package user

import (
	"errors"
	"fmt"
)

//ErrAccountNotFoundById informs that user.Account with the Id was not found in the system.
type ErrAccountNotFoundById struct {
	Id int64
}

func (e ErrAccountNotFoundById) Error() string {
	return fmt.Sprintf("account(id='%d') was not found", e.Id)
}

func IsNotFoundById(err error) bool {
	return errors.As(err, &ErrAccountNotFoundById{})
}

//ErrAccountNotFoundByEmail informs that user.Account with the Email was not found in the system.
type ErrAccountNotFoundByEmail struct {
	Email string
}

func (e ErrAccountNotFoundByEmail) Error() string {
	return fmt.Sprintf("account(email='%s') was not found", e.Email)
}

func IsNotFoundByEmail(err error) bool {
	return errors.As(err, &ErrAccountNotFoundByEmail{})
}

//ErrAccountNotFoundByExternalAccount informs that user.Account with the ExternalAccount was not found.
type ErrAcountNotFoundByExternalAccount struct {
	ExternalAccount ExternalAccount
}

func (e ErrAcountNotFoundByExternalAccount) Error() string {
	return fmt.Sprintf(
		"account with external account(service='%s' id='%s') not found",
		e.ExternalAccount.Service,
		e.ExternalAccount.ExternalId,
	)
}

func IsNotFoundByExternalAccount(err error) bool {
	fmt.Printf("%+v\n", err)
	return errors.As(err, &ErrAcountNotFoundByExternalAccount{})
}

type ErrAccountNotFoundByPhone struct {
	Phone string
}

func (e ErrAccountNotFoundByPhone) Error() string {
	return fmt.Sprintf("account(phone='%s') was not found", e.Phone)
}

func IsNotFoundByPhone(err error) bool {
	return errors.As(err, &ErrAccountNotFoundByPhone{})
}
