package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/auth/vk"
)

func NewAuthController(
	oauth2 vk.Authenticator,
) *fiber.App {
	c := authController{
		vk: oauth2,
	}

	app := fiber.New()
	app.Get("/login", c.Login)
	app.Get("/vk/callback", c.VkOAuthCallback)

	return app
}

type authController struct {
	vk vk.Authenticator
}

func (a authController) Login(c *fiber.Ctx) error {
	return c.Redirect(a.vk.LoginUrl("/"), fiber.StatusSeeOther)
}

func (a authController) VkOAuthCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Empty code")
	}

	_, err := a.vk.Authenticate(c.Context(), vk.OAuth2AccessCodeCredentials{
		ProviderName: "Vk",
		AccessCode:   code,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("OAuthentication error: %+v\n", err))
	}

	c.SendString("Nice")

	return nil
}
