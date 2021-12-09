package config

import (
	"adapteris/config"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrNoAccessCode = errors.New("no access code")
	ErrInvalidProvider = errors.New("no access code")
)

type AuthConfig struct{
	ReceiveAccessCodeUrl string
	Providers []authProviderConfig
}

type AuthApp struct {
	secret []byte
	providers map[string] interface{}
	*fiber.App
}

//NewAuthenticationApplication
func NewAuthenticationApplication(cfg *config.Config) *fiber.App{
	app := AuthApp{
		secret: cfg.Secret,
		App: fiber.New(),
	}
	app.Get("/:provider", Authenticate)
	app.Get("/refresh", RefreshToken)
	return app.App
}

//Authenticate handles OAuth2 access code reception and exchange for the access token of this application.
func Authenticate(c *fiber.Ctx) error {
	if provider := c.Params("provider"); provider == "" {
		c.Status(400)
		return ErrInvalidProvider
	}

	if rawCode := c.Query("code"); rawCode == "" {
		c.Status(400)
		return ErrNoAccessCode
	}

	/*todo: exchange with VK, receive AccessToken, find the userId by ExternalAccount or schedule a creation of new one
	 * in case of absence, generate JWT with userId and roles.
	 **/
	return nil
}

func RefreshToken(c *fiber.Ctx) error {
	return nil
}
