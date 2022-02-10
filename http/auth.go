package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/auth"
	"github.com/sealoftime/adapteris/auth/vk"
)

func NewAuthController(
	auth auth.App,
) *fiber.App {
	api := fiber.New()
	api.Get("/login", login(auth.ConcatLoginUrl))
	api.Get("/vk/callback", signInByVk(auth))

	return api
}

func login(loginUrlWithRedirect func(string) string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Redirect(loginUrlWithRedirect("/"), fiber.StatusSeeOther)
	}
}

func signInByVk(app auth.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Empty code")
		}

		user, err := app.Authenticate(c.Context(), vk.AccessCodeCreds{
			AccessCode: code,
		})

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("OAuthentication error: %+v\n", err))
		}

		return c.SendString("Oh fuck, you must be " + user.ShortName)
	}
}
