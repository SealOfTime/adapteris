package user

import "context"

type Service struct {
	accounts Repository
}

func (s *Service) RetrieveUserById(ctx context.Context, uid int64) (*Account, error) {
	return s.accounts.FindById(ctx, uid)
}
