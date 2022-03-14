//Package school implements domain model for the school designing context.
//i.e. everything that is related to planning the schedule of events and their codependency, but
//not the grading or content of events.
package school

import (
	"context"
	"fmt"
	"time"
)

type School struct {
	Id                 int64
	Name               string
	Visible            bool
	StartDate, EndDate *time.Time
	//Time interval during which new students may join the school.
	RegistrationOpenDate, RegistrationCloseDate *time.Time
	//School Stages - several school steps united by common themes i.e. Soft-Skills, Theory and etc
	Stages []Stage
}

type SchoolRepository interface {
	FindById(ctx context.Context, id int64) (*School, error)
	Save(ctx context.Context, school *School) (*School, error)
}

type SchoolNotFoundByIdErr struct {
	Id int64
}

func (e SchoolNotFoundByIdErr) Error() string {
	return fmt.Sprintf("school with id %d not found", e.Id)
}
