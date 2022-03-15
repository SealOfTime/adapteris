package http

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

type StepHandlers struct {
	*fiber.App
	log          log.Logger
	stepService  *school.StepService
	eventService *school.EventService
}

func NewStepHandlers(
	log log.Logger,
	stepService *school.StepService,
	eventService *school.EventService,
) *StepHandlers {
	app := &StepHandlers{
		App:          fiber.New(),
		log:          log,
		stepService:  stepService,
		eventService: eventService,
	}
	step := app.Group("/:stepId")
	{
		step.Post("/events", app.AddEvent)
	}
	return app
}

type Step struct {
	Id           int64   `json:"id"`
	MustComplete int32   `json:"mustComplete"`
	Events       []Event `json:"events"`
}

func (h *StepHandlers) AddEvent(c *fiber.Ctx) error {
	//Path param
	var (
		stepId int64
	)
	type RequestBody struct {
		Name string `json:"name"`
	}
	type Response struct {
		Event Event `json:"event"`
	}

	var (
		body RequestBody
		res  *school.Event
		ctx  context.Context
		err  error
	)
	if stepId, err = strconv.ParseInt(c.Params("stepId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad stepId",
		)
	}

	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx = c.UserContext()

	if res, err = h.eventService.AddEvent(ctx, school.AddEventRequest{
		StepId: stepId,
		Name:   body.Name,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}
	return c.
		Status(fiber.StatusCreated).
		JSON(Response{
			Event: domainEventToDto(*res),
		})
}

func domainStepToDto(step school.Step) Step {
	res := Step{
		Id:           step.Id,
		MustComplete: step.MustComplete,
		Events:       make([]Event, len(step.Events)),
	}
	for i, event := range step.Events {
		res.Events[i] = domainEventToDto(event)
	}
	return res
}
