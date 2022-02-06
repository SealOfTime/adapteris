package domain

import (
	"context"
	"time"
)

type School struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	Name string
}

type SchoolStorage interface {
	FindById(ctx context.Context, id int64) (*School, error)
	Create(ctx context.Context, user User) (*User, error)
}

// SchoolStage это Этап Школы - несколько событий, связанных тематически. Имеент значение только для наглядности программы школы.
type SchoolStage struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	Name  string
	Color *string
}

// SchoolStep это Шаг Школы - одно событие или несколько событий. В случае нескольких событий выполненность шага определяется
// выполненностью RequiredEvents событий из всех событий (Events) Шага.
type SchoolStep struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	Events         []Event
	RequiredEvents int

	DependsOn *SchoolStep
}
