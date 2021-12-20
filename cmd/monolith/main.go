package main

import (
	"adapteris"
	"adapteris/auth"
	authApi "adapteris/auth/api"
	authStore "adapteris/auth/db"
	"adapteris/config"
	"adapteris/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := adapteris.App{}
	loadConfig(&app)
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
	db := setupDbConnection()
	app.Auth = auth.New(&app.Cfg.Auth, authStore.New(db))
}

func setupDbConnection() *gorm.DB {
	pgDs := util.NewPostgresDS("localhost", "adapteris", "adapteris", "adapteris", 5432)
	db, err := gorm.Open(postgres.Open(pgDs), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open the database connection")
	}
	return db
}

func wrapWithHttpServer(app *adapteris.App) *fiber.App {
	webApp := fiber.New()
	webApp.Mount("/auth", authApi.Wrap(app.Auth))
	return webApp
}

func runHttp(webApp *fiber.App) error {
	return webApp.Listen(":8080")
}
