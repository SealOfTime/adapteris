package adapteris

import (
	"adapteris/auth"
	"adapteris/config"
	"github.com/gofiber/fiber/v2"
)

type App struct{
	Cfg *config.Config
	Auth auth.App
	*fiber.App
}
