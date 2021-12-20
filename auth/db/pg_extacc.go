package db

import (
	"adapteris/auth"
	"context"
	"errors"
	"gorm.io/gorm"
)

type PgExternalAccountRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.ExternalAccountStorage {
	return &PgExternalAccountRepository{db: db}
}

func (p *PgExternalAccountRepository) FindAll(ctx context.Context) ([]auth.ExternalAccount, error) {
	panic("todo")
}

func (p *PgExternalAccountRepository) FindByExternalId(ctx context.Context, extId string) (account auth.ExternalAccount, err error) {
	res := p.db.Where("external_id = ?", extId).Take(&account)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return account, auth.ErrExternalAccountNotFound{ExtId: extId}
	}

	return account, nil
}
