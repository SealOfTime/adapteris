package user

import "context"

type Service struct {
	accounts Repository
}

func NewUserService(accounts Repository) *Service {
	return &Service{
		accounts: accounts,
	}
}

func (s *Service) RetrieveUserById(ctx context.Context, uid int64) (*Account, error) {
	return s.accounts.FindById(ctx, uid)
}
