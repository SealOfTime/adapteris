package setup

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func (a *app) httpServer() *fiber.App {
	app := fiber.New()
	a.setupRoutes(app)
	return app
}

func (a *app) setupRoutes(app *fiber.App) {
	a.mountStaticFiles(app)
	a.mountAPI(app)
	a.setupSPA(app)
}

//mountStaticFiles mounts all the frontend files to the root of app with valid MIME-TYPE
func (a *app) mountStaticFiles(app *fiber.App) {
	requestFileExtensionPattern := regexp.MustCompile(".*\\.([a-z]+)")

	app.Static("/", "./static", fiber.Static{
		Next: func(c *fiber.Ctx) bool {
			// Fix broken mime-types on windows, see: https://github.com/golang/go/issues/32350
			match := requestFileExtensionPattern.FindStringSubmatch(c.OriginalURL())
			if len(match) < 2 {
				return false //don't skip files without extensions
			}

			extension := match[1]
			c.Type(extension)
			return false
		},
	})
}

func (a *app) mountAPI(app *fiber.App) {
	api := app.Group("/api")
	a.mountAuthAPI(api)

	//404 for api calls
	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).
			SendString("This route does not exist")
	})
}

//setupSpa ensures all the redirects neccessary for the proper functioning of same host SPA.
//
// - Any non existent route forwards to /static/index.html.
//
// - /static/index.html redirects to /.
func (a *app) setupSPA(app *fiber.App) {
	// Single-Page Application
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("static/index.html")
	})
}
