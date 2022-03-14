package school

import "context"

const (
	AllEvents int32 = -1
)

type Event struct {
	Id          int64
	Name        string
	Description *string
	Sessions    []EventSession
}

type EventRepository interface {
	FindById(ctx context.Context, id int64) (*Event, error)
	Save(ctx context.Context, event *Event) (*Event, error)
}
