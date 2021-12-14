package main

import (
	"adapteris"
	"adapteris/auth"
	authApi "adapteris/auth/api"
	"adapteris/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := adapteris.App{}
	loadConfig(&app);
	setupDependencies(&app)
	webApi := wrapWithHttpServer(&app)
	if err := runHttp(webApi); err != nil {
		log.Fatalf("Error while running http server: %+v", err)
	}
}

func loadConfig(app *adapteris.App) {
	app.Cfg = config.Read()
}

func setupDependencies(app *adapteris.App) {
	app.Auth = auth.New()
}

func wrapWithHttpServer(app *adapteris.App) (*fiber.App) {
	webApp := fiber.New()
	webApp.Mount("/auth", authApi.Wrap(app.Auth))
	return webApp
}

func runHttp(webApp *fiber.App) (error) {
	return webApp.Listen(":8080");
}
