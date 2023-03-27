package main

import (
	"log"

	"github.com/cn-lxy/music-api/register"
	"github.com/cn-lxy/music-api/tools/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	log.Printf("%v\n", config.Cfg)

	app := fiber.New()

	app.Use(logger.New())
	// app.Use(middleware.VerifyMiddleware)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
		Filter:     jwtFilter,
	}))

	register.Register(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Println("listening")
	log.Fatal(app.Listen(":8000"))
}

// jwt verify filter
func jwtFilter(c *fiber.Ctx) bool {
	path := c.Path()
	skipPath := []string{"/login", "/register"}
	for _, v := range skipPath {
		if path == v {
			return true
		}
	}
	return false
}
