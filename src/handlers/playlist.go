package handlers

import (
	"log"
	"strconv"

	"github.com/cn-lxy/music-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// PlaylistHandler is get user's all playlist hander fiber handler
// the api is "/playlist"
func PlaylistHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := uint64(claims["id"].(float64))
	log.Printf("user id: %d\n", id)
	playlists, err := models.GetPlaylists(id)
	log.Printf("playlists: %v\n", playlists)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "get playlist failed",
		})
	}
	// return user's playlist
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"data": playlists,
	})
}

// CreatePlaylistHandler create a playlist for user
func CreatePlaylistHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// 通过token获取用户ID
	id := uint64(claims["id"].(float64))

	// 解析post json表单
	body := struct {
		Name string `json:"name"`
	}{}
	// 解析失败，返回相应
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "json form error",
		})
	}

	// 创建歌单并写入MySQL数据库并在MogoDB中创建歌单
	pl := models.Playlist{
		CreateUserId: id,
		Name:         body.Name,
	}
	if err := pl.Insert(); err != nil {
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"msg":  "create playlist success.",
	})
}

// DeletePlaylistHandler delete a playlist by playlisy id
func DeletePlaylistHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// 通过token获取用户ID
	id := uint64(claims["id"].(float64))

	pid_string := c.Params("id", "")
	if pid_string == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "query param of id is null",
		})
	}
	pid, err := strconv.ParseUint(pid_string, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "query param of id' value is unallowed",
		})
	}

	// 从MySQL和MongoDB中删除歌单
	pl := models.Playlist{
		Id:           pid,
		CreateUserId: id,
	}
	if err := pl.Delete(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  "delete playlist failed",
		})
	}
	plIds := models.PlaylistSongIds{
		PlaylistId: pid,
	}
	if err := plIds.DeletePlaylistSongIds(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  "delete playlist failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"msg":  "delete playlist success",
	})
}
