package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/domain/user"
	"github.com/sealoftime/adapteris/log"
)

type ProfileHandlers struct {
	*fiber.App
	log      log.Logger
	accounts user.Repository
}

func NewProfileHandlers(
	log log.Logger,
	auth *AuthHandlers,
	accounts user.Repository,
) *ProfileHandlers {
	app := &ProfileHandlers{
		App:      fiber.New(),
		log:      log,
		accounts: accounts,
	}
	app.Get("/my", auth.Authenticated(), app.MyProfile)
	return app
}

type Profile struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`

	FullName  string `json:"fullname"`
	ShortName string `json:"shortname"`

	Email    string `json:"email"`
	Telegram string `json:"tg"`
	Vk       string `json:"vk"`
	Phone    string `json:"phone"`
}

func (h *ProfileHandlers) MyProfile(c *fiber.Ctx) error {
	type Response struct {
		Profile
	}
	var (
		me *user.Account
		ok bool
	)
	if me, ok = c.Locals(UserCtxKey).(*user.Account); !ok {
		panic(fmt.Sprintf("invalid user type: expected *user.Account, got %T", me))
	}

	return c.JSON(Response{
		Profile: profileToDto(me),
	})
}

func profileToDto(user *user.Account) Profile {
	res := Profile{
		Id:        user.Id,
		Role:      string(user.Role),
		ShortName: user.ShortName,
		Email:     user.Email,
		Vk:        user.Vk,
	}
	if user.FullName != nil {
		res.FullName = *user.FullName
	}
	if user.Telegram != nil {
		res.Telegram = *user.Telegram
	}
	if user.PhoneNumber != nil {
		res.Phone = *user.PhoneNumber
	}
	return res
}
