package school

import (
	"context"
	"time"
)

//Stage is a set of school.Steps, that are united by a common theme (i.e. "Soft Skills", "Administrative Theory" etc)
type Stage struct {
	Id                 int64
	Name               string
	Description        string
	StartDate, EndDate *time.Time
	Steps              []Step
}

func (st *Stage) AddStep(step *Step) {
	st.Steps = append(st.Steps, *step)
}

func (stage *Stage) RemoveStep(removed *Step) {
	for i, step := range stage.Steps {
		if step.Id == removed.Id {
			stage.Steps = append(stage.Steps[:i], stage.Steps[i+1:]...)
			return
		}
	}

	panic("removed is not part of Stage's steps")
}

type StageRepository interface {
	FindById(ctx context.Context, id int64) (*Stage, error)
	Save(ctx context.Context, stage *Stage) (*Stage, error)
}
