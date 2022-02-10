package setup

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/auth"
	"github.com/sealoftime/adapteris/auth/vk"
	"github.com/sealoftime/adapteris/http"
)

func (a *app) setupAuth() {
	vkAuth := vk.NewAuthenticator(
		a.cfg.Vk.ClientId,
		a.cfg.Vk.Secret,
		fmt.Sprintf("%s/api/auth/vk/callback", a.cfg.HostURL),
		a.storage.externalAccount,
		a.storage.user,
	)
	a.module.auth = auth.NewApp(vkAuth.LoginUrl, map[string]auth.Authenticator{
		"vk": vkAuth,
	})
}

func (a *app) mountAuthAPI(app fiber.Router) {
	app.Mount("/auth", http.NewAuthController(a.module.auth))
}
