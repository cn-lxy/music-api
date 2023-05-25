package register

import (
	"github.com/cn-lxy/music-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)
	app.Get("/check", handlers.CheckdHandler)

	app.Get("/playlist", handlers.PlaylistHandler)                  // 获取用户所有歌单
	app.Post("/playlist", handlers.CreatePlaylistHandler)           // 创建歌单
	app.Get("/playlist/:id", handlers.PlaylistSongIdsGetIdsHandler) // 获取歌单所有歌曲ID
	app.Put("/playlist/:id", handlers.PlaylistSongIdsAddSong)       // 更新歌曲到歌单（添加或者删除）
	app.Delete("/playlist/:id", handlers.DeletePlaylistHandler)     // 删除歌单
}
