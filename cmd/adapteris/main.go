package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/sealoftime/adapteris/auth/vk"
	"github.com/sealoftime/adapteris/http"
	"github.com/sealoftime/adapteris/mock"
)

var RequestFileExtensionPattern = regexp.MustCompile(".*\\.([a-z]+)")

var VkId = flag.String("vkid", "", "-vkid 18488484")
var VkSecret = flag.String("vksecret", "", "-vksecret okfd8Fmfdjfd838")

func main() {
	flag.Parse()

	if *VkId == "" {
		log.Fatal("No Vk Client ID")
	}

	if *VkSecret == "" {
		log.Fatal("No Vk Client Secret")
	}

	app := fiber.New()
	vkAuth := vk.NewVkAuthenticator(
		*VkId,
		*VkSecret,
		"http://adapteris.test/api/auth/vk/callback",
		&mock.InMemExternalAccountStore{},
		&mock.InMemUserStore{},
	)
	auth := http.NewAuthController(vkAuth)

	app.Static("/", "./static", fiber.Static{
		Next: func(c *fiber.Ctx) bool {
			// Fix broken mime-types on windows, see: https://github.com/golang/go/issues/32350
			match := RequestFileExtensionPattern.FindStringSubmatch(c.OriginalURL())
			if len(match) < 2 {
				return false
			}

			extension := match[1]
			c.Type(extension)
			return false
		},
	})

	api := app.Group("/api")
	{
		api.Mount("/auth", auth)
		api.Get("/hello", func(c *fiber.Ctx) error {
			name := c.Query("name")
			return c.SendString(fmt.Sprintf("Hello, %s!", name))
		})

		//404 for api calls
		api.Use(func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).
				SendString("This route does not exist")
		})
	}

	// Single-Page Application
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("static/index.html")
	})

	if err := app.Listen(":8080"); err != nil {
		fmt.Println(err)
	}
}
