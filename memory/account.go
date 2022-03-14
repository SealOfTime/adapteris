package memory

import (
	"context"

	"github.com/sealoftime/adapteris/domain/user"
)

type AccountStore struct {
	id       int64
	accounts []*user.Account
}

func (a *AccountStore) Save(ctx context.Context, acc user.Account) (*user.Account, error) {
	if acc.Id == 0 {
		return a.create(ctx, acc)
	}

	for _, ac := range a.accounts {
		if ac.Id == acc.Id {
			ac.Role = acc.Role

			ac.ShortName = acc.ShortName
			ac.FullName = acc.FullName

			ac.Email = acc.Email
			ac.Telegram = acc.Telegram
			ac.Vk = acc.Vk
			ac.PhoneNumber = acc.PhoneNumber

			ac.ExternalAccounts = acc.ExternalAccounts
			return ac, nil
		}
	}

	return nil, user.ErrAccountNotFoundById{Id: acc.Id}
}

func (a *AccountStore) create(ctx context.Context, acc user.Account) (*user.Account, error) {
	a.id++
	acc.Id = a.id
	a.accounts = append(a.accounts, &acc)
	return a.accounts[len(a.accounts)-1], nil
}
