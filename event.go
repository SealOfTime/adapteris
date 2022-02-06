package domain

import "time"

// Event это мероприятие школы.
type Event struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	Name        string
	Description string

	GradeCriterias []GradeCriteria
	Sessions       []EventSession
}

// EventSession это Проведение Мероприятия школы: одно мероприятие может проводиться несколько раз, чтобы участники могли выбрать
// удобное для них время и место, и это как раз объединение события, времени, места и тех, кто будет это всё проводить.
type EventSession struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	Event *Event
	//Дата и время проведения мероприятия
	Time       time.Time
	Place      string
	Organizers []User
}
