package school

import (
	"context"
	"fmt"

	"github.com/sealoftime/adapteris/log"
)

type EventService struct {
	events EventRepository
	steps  StepRepository
}

func NewEventService(
	log log.Logger,
	steps StepRepository,
	events EventRepository,
) *EventService {
	return &EventService{
		steps:  steps,
		events: events,
	}
}

type AddEventRequest struct {
	StepId int64
	Name   string
}

func (s *EventService) AddEvent(ctx context.Context, req AddEventRequest) (*Event, error) {
	var (
		step  *Step
		saved *Step
		err   error
	)
	if step, err = s.steps.FindById(ctx, req.StepId); err != nil {
		return nil, fmt.Errorf("couldn't find step for adding event: %w", err)
	}

	step.Events = append(step.Events, Event{
		Name: req.Name,
	})

	if saved, err = s.steps.Save(ctx, step); err != nil {
		return nil, fmt.Errorf("couldn't save step after adding event: %w", err)
	}
	for i, original := range step.Events {
		if saved.Events[i].Id != original.Id {
			return &saved.Events[i], nil
		}
	}

	panic("no difference, doesn't matter")
}

type EventRenameRequest struct {
	EventId int64
	NewName string
}

func (s *EventService) Rename(ctx context.Context, req EventRenameRequest) error {
	event, err := s.events.FindById(ctx, req.EventId)
	if err != nil {
		return err
	}

	event.Name = req.NewName

	if _, err := s.events.Save(ctx, event); err != nil {
		return err
	}

	return nil
}

type EventChangeDescriptionRequest struct {
	EventId        int64
	NewDescription string
}

func (s *EventService) ChangeDescription(ctx context.Context, req EventChangeDescriptionRequest) error {
	event, err := s.events.FindById(ctx, req.EventId)
	if err != nil {
		return err
	}

	event.Description = &req.NewDescription

	if _, err := s.events.Save(ctx, event); err != nil {
		return err
	}

	return nil
}
