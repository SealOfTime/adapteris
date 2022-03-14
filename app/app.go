package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/log"
)

type App struct {
	Log     log.Logger
	Config  Config
	Storage Storages
	Service Services
	Routes
	Http *fiber.App
}

func New(cfg Config) App {
	a := App{Config: cfg}
	a.initLogger()
	a.initStorage()
	a.initServices()
	a.initRoutes()
	a.initHttpserver()
	return a
}

func (a *App) Start() error {
	return a.Http.Listen(fmt.Sprintf(":%d", a.Config.Port))
}
