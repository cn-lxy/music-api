package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/cn-lxy/music-api/middleware"
	"github.com/cn-lxy/music-api/register"
	"github.com/cn-lxy/music-api/tools"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Printf("%v\n", tools.Cfg)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.VerifyMiddleware)

	register.Register(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Println("listening")
	log.Fatal(app.Listen(":8000"))
}
