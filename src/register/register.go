package register

import (
	"github.com/cn-lxy/music-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)
	app.Get("/check", handlers.CheckdHandler)

	app.Get("/playlist", handlers.PlaylistHandler)              // 获取用户所有歌单
	app.Post("/playlist", handlers.CreatePlaylistHandler)       // 创建歌单
	app.Delete("/playlist/:id", handlers.DeletePlaylistHandler) // 删除歌单
}
