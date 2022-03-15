package http

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
)

type EventHandlers struct {
	*fiber.App
	log          log.Logger
	eventService *school.EventService
}

func NewEventHandlers(
	log log.Logger,
	eventService *school.EventService,
) *EventHandlers {
	app := &EventHandlers{
		App:          fiber.New(),
		log:          log,
		eventService: eventService,
	}
	event := app.Group("/:eventId")
	{
		event.Put("/name", app.Rename)
		event.Put("/description", app.ChangeDescription)
	}
	return app
}

type Event struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//CRUD
func (h *EventHandlers) Rename(c *fiber.Ctx) error {
	//Path param.
	var (
		eventId int64
	)
	type RequestBody struct {
		NewName string `json:"newName"`
	}

	var err error
	if eventId, err = strconv.ParseInt(c.Params("eventId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad eventId",
		)
	}
	var body RequestBody
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx := c.UserContext()

	if err = h.eventService.Rename(ctx, school.EventRenameRequest{
		EventId: eventId,
		NewName: body.NewName,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *EventHandlers) ChangeDescription(c *fiber.Ctx) error {
	//Path param.
	var (
		eventId int64
	)
	type RequestBody struct {
		NewDescription string `json:"newDescription"`
	}

	var err error
	if eventId, err = strconv.ParseInt(c.Params("eventId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad eventId",
		)
	}
	var body RequestBody
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}

	ctx := c.UserContext()

	if err = h.eventService.ChangeDescription(ctx, school.EventChangeDescriptionRequest{
		EventId:        eventId,
		NewDescription: body.NewDescription,
	}); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func domainEventToDto(event school.Event) Event {
	res := Event{
		Id:   event.Id,
		Name: event.Name,
	}
	if event.Description != nil {
		res.Description = *event.Description
	}
	return res
}
