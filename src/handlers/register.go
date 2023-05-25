package handlers

import (
	"log"

	"github.com/cn-lxy/music-api/models"
	"github.com/cn-lxy/music-api/tools/jwt"
	"github.com/gofiber/fiber/v2"
)

// RegisterHandler user register
func RegisterHandler(c *fiber.Ctx) error {
	// create new user instance
	user := &models.User{}
	// parse body into user model
	if err := c.BodyParser(user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  err.Error(),
		})
	}
	// insert user into database
	if id, err := user.Insert(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  err.Error(),
		})
	} else {
		// 创建歌单并写入MySQL数据库并在MogoDB中创建歌单
		pl := models.Playlist{
			CreateUserId: uint64(id),
			Name:         "__LIKE__",
		}
		if _, err := pl.Insert(); err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": fiber.StatusInternalServerError,
				"msg":  "create playlist fail.",
			})
		}
		plIds := models.PlaylistSongIds{
			PlaylistId: pl.Id,
		}
		log.Println(plIds)
		if err := plIds.CreatePlaylistSongIds(); err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": fiber.StatusInternalServerError,
				"msg":  "create playlist fail.",
			})
		}
	}

	// generate JWT token
	token, err := jwt.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  err.Error(),
		})
	}

	// return JWT token and code
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code": fiber.StatusCreated,
		"msg":  "success",
		"data": fiber.Map{
			"token": token,
		},
	})
}
