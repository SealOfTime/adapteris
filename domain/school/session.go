package school

import (
	"context"
	"time"

	"github.com/sealoftime/adapteris/domain/user"
)

type EventSession struct {
	Id              int64
	Time            time.Time
	Place           string
	MaxParticipants int32
	Organizers      []user.Account
	Participants    []Participation
}

type EventSessionRepository interface {
	FindById(ctx context.Context, id int64) (*EventSession, error)
	Save(ctx context.Context, session *EventSession) (*EventSession, error)
}
