package school

import (
	"context"

	"github.com/sealoftime/adapteris/domain/user"
)

type Participation struct {
	Id          int64
	Participant user.Account
	Passed      bool
}

type ParticipationRepository interface {
	FindById(ctx context.Context, id int64) (*Participation, error)
	Save(ctx context.Context, saved *Participation) (*Participation, error)
}
