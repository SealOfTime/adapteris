// Package api provides API for the authentication logic.
package api

import (
	"adapteris/auth"
	"errors"
	"github.com/gofiber/fiber/v2"
)

// Wrap function builds fiber application proxing the incoming requests to the provided auth.App instance.
func Wrap(app auth.App) *fiber.App {
	api := fiber.New()
	api.Get("/:provider", authenticate(app))
	api.Get("/refresh", refreshToken(app))
	return api
}

var (
	ErrNoAccessCode    = errors.New("no access code")
	ErrInvalidProvider = errors.New("no provider")
)

// authenticate initialises new handler in the context of auth.App, that will confirm the identity of the user and
// respond with JWT Token for further authentications.
func authenticate(app auth.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if provider := c.Params("provider"); provider == "" {
			c.Status(400)
			return ErrInvalidProvider
		}

		if rawCode := c.Query("code"); rawCode == "" {
			c.Status(400)
			return ErrNoAccessCode
		}

		return nil
	}
}

func refreshToken(app auth.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
