package school

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	schools SchoolRepository
}

func NewService(schoolRepo SchoolRepository) *Service {
	return &Service{
		schools: schoolRepo,
	}
}

func (s *Service) GetSchool(ctx context.Context, id int64) (*School, error) {
	return s.schools.FindById(ctx, id)
}

type CreateRequest struct {
	Name                  string    `json:"name"`
	StartDate             time.Time `json:"start"`
	EndDate               time.Time `json:"end"`
	RegistrationOpenDate  time.Time `json:"registerStart"`
	RegistrationCloseDate time.Time `json:"registerEnd"`
}

func (s *Service) Create(ctx context.Context, req CreateRequest) (*School, error) {
	if req.EndDate.Before(req.StartDate) {
		return nil, fmt.Errorf("school's end is earlier than its start")
	}

	if req.RegistrationCloseDate.Before(req.RegistrationOpenDate) {
		return nil, fmt.Errorf("school's registration close date is earlier than its open date")
	}

	if req.Name == "" {
		if req.StartDate.IsZero() {
			return nil, fmt.Errorf("school name must not be empty, when its start date is unclear")
		}

		req.Name = fmt.Sprintf("Школа %d года", req.StartDate.Year())
	}

	return s.schools.Save(ctx, &School{
		Name:                  req.Name,
		StartDate:             &req.StartDate,
		EndDate:               &req.EndDate,
		RegistrationOpenDate:  &req.RegistrationOpenDate,
		RegistrationCloseDate: &req.RegistrationCloseDate,
	})
}

type RenameRequest struct {
	SchoolId int64
	NewName  string
}

func (s *Service) Rename(ctx context.Context, req RenameRequest) error {
	school, err := s.schools.FindById(ctx, req.SchoolId)
	if err != nil {
		return err
	}

	school.Name = req.NewName

	if _, err := s.schools.Save(ctx, school); err != nil {
		return err
	}

	return nil
}

type ScheduleSchoolReq struct {
	SchoolId  int64
	StartDate time.Time
	EndDate   time.Time
}

func (s *Service) ScheduleSchool(ctx context.Context, req ScheduleSchoolReq) error {
	if req.EndDate.Before(req.StartDate) {
		return fmt.Errorf("school ends earlier than it ends")
	}

	school, err := s.schools.FindById(ctx, req.SchoolId)
	if err != nil {
		return err
	}

	school.StartDate = &req.StartDate
	school.EndDate = &req.EndDate

	if _, err := s.schools.Save(ctx, school); err != nil {
		return err
	}

	return nil
}

type ScheduleRegistrationReq struct {
	SchoolId  int64
	StartDate time.Time
	EndDate   time.Time
}

func (s *Service) ScheduleRegistration(ctx context.Context, req ScheduleRegistrationReq) error {
	if req.EndDate.Before(req.StartDate) {
		return fmt.Errorf("school registration end earlier than it starts")
	}

	school, err := s.schools.FindById(ctx, req.SchoolId)
	if err != nil {
		return err
	}

	school.RegistrationOpenDate = &req.StartDate
	school.RegistrationCloseDate = &req.EndDate

	if _, err := s.schools.Save(ctx, school); err != nil {
		return err
	}

	return nil
}
