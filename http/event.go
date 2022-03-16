package http

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/school"
	"github.com/sealoftime/adapteris/log"
	"github.com/sealoftime/adapteris/pgx"
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
		event.Get("/", app.GetEvent)
		event.Put("/name", app.Rename)
		event.Put("/description", app.ChangeDescription)
		event.Post("/sessions", app.AddSession)
	}
	return app
}

type Event struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Sessions    []EventSession `json:"sessions"`
}

func (h *EventHandlers) AddSession(c *fiber.Ctx) error {
	//Path param.
	var (
		eventId int64
	)
	type RequestBody struct {
		Session EventSession `json:"session"`
	}
	var (
		ctx  = c.UserContext()
		body RequestBody
		err  error
	)
	if eventId, err = strconv.ParseInt(c.Params("eventId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad eventId",
		)
	}
	if err = c.BodyParser(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			fmt.Sprintf("malformed json: %+v", err),
		)
	}
	tx := pgx.TxFromCtx(ctx)
	if _, err = tx.Exec(ctx,
		"INSERT INTO event_session (place, datetime, max_participants, event_id) VALUES ($1, $2, $3, $4)",
		body.Session.Place, body.Session.DateTime, 0, eventId,
	); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *EventHandlers) GetEvent(c *fiber.Ctx) error {
	//Path param.
	var (
		eventId int64
	)
	type Response struct {
		Event
	}

	var (
		ctx   = c.UserContext()
		event *school.Event
		err   error
	)
	if eventId, err = strconv.ParseInt(c.Params("eventId"), 10, 64); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"bad eventId",
		)
	}
	if event, err = h.eventService.Get(ctx, eventId); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unexpected error: %+v", err),
		)
	}
	return c.JSON(Response{
		Event: domainEventToDto(*event),
	})
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
		Id:       event.Id,
		Name:     event.Name,
		Sessions: make([]EventSession, len(event.Sessions)),
	}
	for i, ses := range event.Sessions {
		res.Sessions[i] = domainSessionToDto(ses)
	}

	if event.Description != nil {
		res.Description = *event.Description
	}
	return res
}
