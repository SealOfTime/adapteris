package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sealoftime/adapteris/domain/user"
	"github.com/sealoftime/adapteris/integration/vk"
)

const (
	UserCtxKey       = "USER"
	UserIdSessionKey = "USER_ID"
)

type AuthHandlers struct {
	*fiber.App
	accountService *user.Service
	sessionService *session.Store
	authService    *user.AuthService
	vkService      *vk.Service
}

func NewAuthController(
	session *session.Store,
	accountsService *user.Service,
	authService *user.AuthService,
	vkService *vk.Service,
) AuthHandlers {
	app := AuthHandlers{
		App:            fiber.New(),
		sessionService: session,
		accountService: accountsService,
		authService:    authService,
		vkService:      vkService,
	}
	app.Get("/login", loginWithVk(vkService))
	app.Get("/vk/callback", app.vkCallback)
	return app
}

func loginWithVk(vkService *vk.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Redirect(vkService.GetLoginUrl(c.UserContext(), "/"), fiber.StatusSeeOther)
	}
}

// 1. Exchange code for token
// 2. Extract ExternalId from token
// 3. Find User by ExternalId
// 3.1 If user is not found extract default parameters from access token and an external service
// 3.2 Create a new user based on the data
// 4. Form a JWT
func (a *AuthHandlers) vkCallback(c *fiber.Ctx) error {
	ctx := c.UserContext()

	code := c.Query("code")
	if code == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Empty code")
	}

	u, err := a.vkService.LoginByAccessCode(ctx, code)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("authentication by vk failed: %+v", err),
		)
	}

	return a.authenticate(c, u.Id)
}

func (a *AuthHandlers) authenticate(c *fiber.Ctx, uid int64) error {
	s, err := a.sessionService.Get(c)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unable to create session: %+v", err),
		)
	}

	s.Set(UserIdSessionKey, true)

	if err := s.Save(); err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			fmt.Sprintf("unable to save session: %+v", err),
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (a *AuthHandlers) Authenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s, err := a.sessionService.Get(c)
		if err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				SendString(fmt.Sprintf("unable to acquire session: %+v", err))
		}

		uidRaw := s.Get(UserIdSessionKey)
		if uidRaw == nil {
			return c.
				SendStatus(fiber.StatusUnauthorized)
		}

		u, err := a.accountService.RetrieveUserById(c.UserContext(), uidRaw.(int64))
		if err != nil {
			return c.
				Status(fiber.StatusInternalServerError).
				SendString(fmt.Sprintf("session contains invalid user: %+v", err))
		}

		c.Locals(UserCtxKey, u)
		return nil
	}
}

func (a AuthHandlers) Authorized(allowedRole user.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, _ := c.Locals(UserCtxKey).(*user.Account)
		if u == nil {
			return c.
				Status(fiber.StatusUnauthorized).
				SendString("Unauthenticated")
		}

		switch {
		case u.Role == user.ADMIN:
			return nil
		case u.Role == allowedRole:
			return nil
		}

		return c.
			Status(fiber.StatusForbidden).
			SendString("Your role does not satisfy the requirement")
	}
}
