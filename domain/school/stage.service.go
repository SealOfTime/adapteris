package school

import (
	"context"
	"time"

	"github.com/sealoftime/adapteris/log"
)

type StageService struct {
	schools SchoolRepository
	stages  StageRepository
	log     log.Logger
}

func NewStageService(
	log log.Logger,
	schoolStore SchoolRepository,
) *StageService {
	return &StageService{
		schools: schoolStore,
		log:     log,
	}
}

type AddStageRequest struct {
	SchoolId    int64
	Name        string
	Description string
	Start       *time.Time
	End         *time.Time
}

func (s *StageService) AddStage(ctx context.Context, req AddStageRequest) (*Stage, error) {
	school, err := s.schools.FindById(ctx, req.SchoolId)
	if err != nil {
		return nil, err
	}

	school.Stages = append(school.Stages, Stage{
		Name:        req.Name,
		Description: req.Description,
		StartDate:   req.Start,
		EndDate:     req.End,
	})

	dbSchool, err := s.schools.Save(ctx, school)
	if err != nil {
		return nil, err
	}

	for _, stage := range dbSchool.Stages {
		if stage.Name == req.Name {
			return &stage, nil
		}
	}

	panic("stage name changed after saving")
}

type StageRenameRequest struct {
	StageId int64
	NewName string
}

func (s *StageService) Rename(ctx context.Context, req StageRenameRequest) error {
	stage, err := s.stages.FindById(ctx, req.StageId)
	if err != nil {
		return err
	}

	stage.Name = req.NewName

	if _, err := s.stages.Save(ctx, stage); err != nil {
		return err
	}

	return nil
}

type StageChangeDescriptionRequest struct {
	StageId        int64
	NewDescription string
}

func (s *StageService) ChangeDescription(ctx context.Context, req StageChangeDescriptionRequest) error {
	stage, err := s.stages.FindById(ctx, req.StageId)
	if err != nil {
		return err
	}

	if stage.Id == req.StageId {
		stage.Description = req.NewDescription
	}

	if _, err := s.stages.Save(ctx, stage); err != nil {
		return err
	}

	return nil
}

type SchedulStageRequest struct {
	StageId int64
	Start   time.Time
	End     time.Time
}

func (s *StageService) ScheduleStage(ctx context.Context, req SchedulStageRequest) error {
	stage, err := s.stages.FindById(ctx, req.StageId)
	if err != nil {
		return err
	}

	if stage.Id == req.StageId {
		stage.StartDate = &req.Start
		stage.EndDate = &req.End
	}

	if _, err := s.stages.Save(ctx, stage); err != nil {
		return err
	}

	return nil
}
