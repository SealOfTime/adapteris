package domain

import "time"

type gradeCriteriaType int

const (
	UNKNOWN gradeCriteriaType = iota
	BooleanCriteria
	DecimalCriteria
	TextCriteria
)

type GradeCriteria struct {
	Id        GradeCriteriaId
	CreatedAt time.Time
	EditedAt  time.Time

	Name string
	Type gradeCriteriaType
}

type GradeCriteriaId = int64

type EventResults struct {
	Id        int64
	CreatedAt time.Time
	EditedAt  time.Time

	User   *User
	Event  *Event
	Passed BooleanGrade

	Grades map[GradeCriteriaId]Grade
}

type Grade interface{ Empty() bool }

type BooleanGrade struct{ bool, empty bool }
type TextGrade string
type DecimalGrade struct{ int64, empty bool }

func (g BooleanGrade) Empty() bool { return g.empty }
func (g TextGrade) Empty() bool    { return g == "" }
func (g DecimalGrade) Empty() bool { return g.empty }
