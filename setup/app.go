package setup

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	domain "github.com/sealoftime/adapteris"
	"github.com/sealoftime/adapteris/auth"
)

type app struct {
	cfg     config
	http    *fiber.App
	module  modules
	storage storages
}

type modules struct {
	auth auth.App
}

type storages struct {
	user            domain.UserStorage
	externalAccount domain.ExternalAccountStorage
}

func App() app {
	var a app
	a.setupConfig()
	a.setupModules()
	a.httpServer()
	return a
}

func (a *app) Start() error {
	return a.http.Listen(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *app) setupConfig() {
	var cfg config
	cfg.parseFlags()
	cfg.parseJson(loadJsonConfig(cfg.ConfigPath))
}

func (a *app) setupModules() {
	a.setupAuth()
}
