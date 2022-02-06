package mock

import (
	"context"

	domain "github.com/sealoftime/adapteris"
)

type InMemExternalAccountStore struct {
	storage []domain.ExternalAccount
}

func (s *InMemExternalAccountStore) FindByExternalId(ctx context.Context, id string) (*domain.ExternalAccount, error) {
	for _, e := range s.storage {
		if e.ExternalId == id {
			return &e, nil
		}
	}

	return nil, domain.ErrExternalAccountNotFound{ExtId: id}
}
