package http

import (
	"adapteris/config"
	"errors"
	"github.com/gofiber/fiber/v2"
)
var (
	ErrNoAccessCode = errors.New("no access code")
)

type AuthService struct {
	secret []byte
	*fiber.App
}

func NewAuthenticationService(cfg *config.Config) *fiber.App{
	app := AuthService{cfg.Secret,fiber.New()}
	app.Get("/refresh", RefreshToken)
	return app.App
}

//Authenticate handles OAuth2 access code reception and exchange for the access token of this application.
func Authenticate(c *fiber.Ctx) error {
	rawCode := c.Query("code")
	if rawCode == "" {
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


