package user

import (
	"context"
)

type AuthService struct {
	accounts Repository
}

func NewAuthService(accounts Repository) *AuthService {
	return &AuthService{
		accounts: accounts,
	}
}

// LoginByExternalAccount logs user in by the external account.
//
//errors: ErrAcountNotFoundByExternalAccount
func (a *AuthService) LoginByExternalAccount(
	ctx context.Context,
	eac ExternalAccount,
) (*Account, error) {
	acc, err := a.accounts.FindByExternalAccount(ctx, eac)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

//RegisterUserByExternalAccount registers a new account or attaches the external account
//to the existing account found by email.
func (a *AuthService) RegisterUserByExternalAccount(
	ctx context.Context,
	defaults Account,
	eac ExternalAccount,
) (*Account, error) {
	account, err := a.accounts.FindByEmail(ctx, defaults.Email)
	if err != nil && !IsNotFoundByEmail(err) {
		return nil, err
	}

	if IsNotFoundByEmail(err) {
		account = &defaults
	}

	account.LinkExternalAccount(eac)
	return a.accounts.Save(ctx, *account)
}
