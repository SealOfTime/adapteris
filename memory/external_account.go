package memory

import (
	"context"

	domain "github.com/sealoftime/adapteris"
)

type ExternalAccountStore struct {
	storage []domain.ExternalAccount
}

func (s *ExternalAccountStore) FindByExternalId(ctx context.Context, id string) (*domain.ExternalAccount, error) {
	for _, e := range s.storage {
		if e.ExternalId == id {
			return &e, nil
		}
	}

	return nil, domain.ErrExternalAccountNotFound{ExtId: id}
}
