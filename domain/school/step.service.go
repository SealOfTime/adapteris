package school

import (
	"context"
	"fmt"

	"github.com/sealoftime/adapteris/log"
)

type StepService struct {
	stages StageRepository
	steps  StepRepository
	log    log.Logger
}

func NewStepsService(
	log log.Logger,
	stages StageRepository,
	steps StepRepository,
) *StepService {
	return &StepService{
		steps:  steps,
		stages: stages,
		log:    log,
	}
}

type AddStepRequest struct {
	StageId int64
}

func (s *StepService) AddStep(ctx context.Context, req AddStepRequest) (*Step, error) {
	stage, err := s.stages.FindById(ctx, req.StageId)
	if err != nil {
		return nil, err
	}

	stage.AddStep(&Step{})

	dbStage, err := s.stages.Save(ctx, stage)
	if err != nil {
		return nil, err
	}

	for i, step := range dbStage.Steps {
		if stage.Steps[i].Id != step.Id {
			return &step, nil
		}
	}
	panic("no mismatching steps found.")
}

type MoveStepRequest struct {
	FromStageId int64
	ToStageId   int64
	StepId      int64
}

func (s *StepService) MoveStep(ctx context.Context, req MoveStepRequest) error {
	from, err := s.stages.FindById(ctx, req.FromStageId)
	if err != nil {
		return fmt.Errorf("error finding stage to move step from: %w", err)
	}
	to, err := s.stages.FindById(ctx, req.ToStageId)
	if err != nil {
		return fmt.Errorf("error finding stage to move step to: %w", err)
	}
	step, err := s.steps.FindById(ctx, req.StepId)
	if err != nil {
		return fmt.Errorf("error finding step to move: %w", err)
	}

	from.RemoveStep(step)
	to.AddStep(step)

	if _, err := s.stages.Save(ctx, to); err != nil {
		return fmt.Errorf("error finding stage, from which step was moved: %w", err)
	}

	if _, err := s.stages.Save(ctx, from); err != nil {
		return fmt.Errorf("error finding stage, to which step was moved: %w", err)
	}

	return nil
}

type TweakMustCompleteRequest struct {
	StepId          int64
	NewMustComplete int32
}

func (s *StepService) TweakMustComplete(ctx context.Context, req TweakMustCompleteRequest) error {
	step, err := s.steps.FindById(ctx, req.StepId)
	if err != nil {
		return fmt.Errorf("error finding step to tweak MustComplete: %w", err)
	}

	if req.NewMustComplete < 0 && req.NewMustComplete != -1 {
		return fmt.Errorf("MustComplete must be either a non-negative or -1")
	}

	step.MustComplete = req.NewMustComplete

	if _, err := s.steps.Save(ctx, step); err != nil {
		return fmt.Errorf("error saving step after tweaking MustComplete: %w", err)
	}

	return nil
}
