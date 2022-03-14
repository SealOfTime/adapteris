package user

import (
	"context"
	"fmt"
	"time"
)

const (
	DEFAULT Role = "USER"
	ADMIN   Role = "ADMIN"
)

type Role string

type Account struct {
	Id           int64
	RegisteredAt time.Time
	Role         Role

	FullName  *string
	ShortName string `validate:"required"`

	Email       string `validate:"required,email"`
	Telegram    *string
	Vk          string `validate:"required,alphanum"`
	PhoneNumber *string

	ExternalAccounts []ExternalAccount
}

func NewAccount(
	shortName,
	email,
	vk string,
) (acc Account, errs []error) {
	if shortName == "" {
		errs = append(errs, fmt.Errorf("shortname must not be empty"))
	}
	if email == "" {
		errs = append(errs, fmt.Errorf("email must not be empty"))
	}
	if vk == "" {
		errs = append(errs, fmt.Errorf("vk link must not be empty"))
	}
	if errs != nil {
		return
	}

	return Account{
		ShortName:    shortName,
		Email:        email,
		Vk:           vk,
		Role:         DEFAULT,
		RegisteredAt: time.Now(),
	}, nil
}

type ExternalAccount struct {
	Id         int64
	Service    string
	ExternalId string
}

func (a *Account) LinkExternalAccount(eac ExternalAccount) {
	a.ExternalAccounts = append(a.ExternalAccounts, eac)
}

type Repository interface {
	//FindByExternalAccount returns an Account found by ExternalAccount.
	//
	//errors: ErrAccountNotFoundByExternalAccount
	FindByExternalAccount(ctx context.Context, eac ExternalAccount) (*Account, error)

	//FindById returns an Account found by Account Id
	//
	//errors: ErrAccountNotFoundById
	FindById(ctx context.Context, id int64) (*Account, error)

	//FindByEmail returns an Account found by Account email.
	//
	//errors: ErrAccountNotFoundByEmail
	FindByEmail(ctx context.Context, email string) (*Account, error)

	Save(ctx context.Context, acc Account) (*Account, error)
}
