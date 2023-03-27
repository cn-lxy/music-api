package register

import (
	"github.com/cn-lxy/music-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)

	app.Get("/playlist", handlers.PlaylistHandler)
}
