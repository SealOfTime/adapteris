package api


import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)
const (
	USER = "user"
)

func Authorized(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		//SuccessHandler:           nil,
		//ErrorHandler:             nil,
		SigningKey:               secret,
		//todo: consider for soa deployment
		ContextKey: USER,
		Claims:     nil,
	})
}
