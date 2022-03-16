package app

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sealoftime/adapteris/http"
)

type Routes struct {
	auth          *http.AuthHandlers
	school        *http.SchoolHandlers
	stage         *http.StageHandlers
	step          *http.StepHandlers
	event         *http.EventHandlers
	profile       *http.ProfileHandlers
	participation *http.ParticipationHandlers
}

func (a *App) initRoutes() {
	a.Routes.auth = http.NewAuthController(
		session.New(),
		a.Config.HostURL,
		a.Service.User,
		a.Service.Auth,
		a.Service.Integration.Vk,
	)
	a.Routes.school = http.NewSchoolHandlers(
		a.Service.School,
		a.Service.Stage,
	)
	a.Routes.stage = http.NewStageHandlers(
		a.Log,
		a.Service.Stage,
		a.Service.Step,
	)
	a.Routes.step = http.NewStepHandlers(
		a.Log,
		a.Service.Step,
		a.Service.Event,
	)
	a.Routes.event = http.NewEventHandlers(
		a.Log,
		a.Service.Event,
	)
	a.Routes.profile = http.NewProfileHandlers(
		a.Log,
		a.Routes.auth,      //todo: bad
		a.Storage.accounts, //todo: very bad
	)
	a.Routes.participation = http.NewParticipationHandlers(
		a.Routes.auth,
		a.Storage.participations,
	)
}

func (a *App) initHttpserver() {
	a.Http = fiber.New()
	a.Http.Use(logger.New())
	a.Http.Use(http.PgxTransactional(a.Storage.connPool, a.Log))
	a.mountRoutes(a.Http)
}

func (a *App) mountRoutes(root *fiber.App) {
	mountStaticFiles(root)
	a.mountAPI(root)
	a.setupSPA(root)
}

//mountStaticFiles mounts all the frontend files to the root of app with valid MIME-TYPE
func mountStaticFiles(app fiber.Router) {
	requestFileExtensionPattern := regexp.MustCompile(`.*\.([a-z]+)`)

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

func (a *App) mountAPI(route fiber.Router) {
	api := route.Group("/api")
	api.Mount("/auth", a.Routes.auth.App)
	api.Mount("/school", a.Routes.school.App)
	api.Mount("/stage", a.Routes.stage.App)
	api.Mount("/step", a.Routes.step.App)
	api.Mount("/event", a.Routes.event.App)
	api.Mount("/profile", a.Routes.profile.App)
	api.Mount("/participations", a.Routes.participation.App)
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
func (a *App) setupSPA(app *fiber.App) {
	// Single-Page Application
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("static/index.html")
	})
}
