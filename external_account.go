package domain

import (
	"context"
	"fmt"
)

type ExternalAccount struct {
	Id         int64
	User       *User
	Service    string
	ExternalId string
}

type ExternalAccountStorage interface {
	FindByExternalId(ctx context.Context, id string) (*ExternalAccount, error)
}

//ErrExternalAccountNotFound signals that the ExternalAccount with this ExtId was not found
type ErrExternalAccountNotFound struct {
	ExtId string
}

func (e ErrExternalAccountNotFound) Error() string {
	return fmt.Sprintf("external account with external id '%s' not found", e.ExtId)
}
