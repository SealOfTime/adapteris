package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/log"
)

type ParticipantHandlers struct {
	*fiber.App
}

func NewParticipantHandlers(
	log log.Logger,
	auth *AuthHandlers,
) *ParticipantHandlers {
	app := &ParticipantHandlers{
		App: fiber.New(),
	}
	app.Use(auth.Authenticated())
	app.Get("/school", app.GetSchool)
	app.Get("/results", app.GetSchool)
	return app
}

func (h *ParticipantHandlers) GetSchool(c *fiber.Ctx) error {
	panic("todo")
}
