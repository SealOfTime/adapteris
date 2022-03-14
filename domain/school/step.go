package school

import "context"

//Step is several school.Events united by "complete X of N" or "complete all of N" relationships,
//representing a milestone in school's course.
type Step struct {
	Id int64
	//MustComplete is an amount of events out of all events of this step to be compeleted
	//for the whole step to be considered completed.
	// 0 for optional step.
	// -1 for all.
	MustComplete int32
	Events       []Event
}

type StepRepository interface {
	FindById(ctx context.Context, id int64) (*Step, error)
	Save(ctx context.Context, step *Step) (*Step, error)
}
