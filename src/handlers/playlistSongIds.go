package handlers

import (
	"log"
	"strconv"

	"github.com/cn-lxy/music-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func PlaylistSongIdsGetIdsHandler(c *fiber.Ctx) error {
	// 通过token获取用户ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := uint64(claims["id"].(float64))
	log.Printf("user id: %d\n", id)

	// 获取路由中的ID
	pid, err := strconv.ParseUint(c.Params("id", "0"), 10, 64)
	if err != nil || pid == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "id error",
		})
	}

	// 获取歌单所有歌曲ID
	ps := models.Playlist{
		Id:           pid,
		CreateUserId: id,
	}
	if err := ps.Get(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  "get playlist failed",
		})
	}
	psIds := models.PlaylistSongIds{
		PlaylistId: pid,
	}
	if err := psIds.GetAllSong(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  "get playlist song ids failed",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"msg":  "success",
		"data": fiber.Map{
			"playlist": ps,
			"songIds":  psIds.SongIds,
		},
	})
}

// PlaoyListSongIdsAddSong 添加歌曲到歌单中的操作。 ?type=add为添加，?type=del为删除
func PlaylistSongIdsUpdate(c *fiber.Ctx) error {
	// 通过token获取用户ID
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := uint64(claims["id"].(float64))
	log.Printf("user id: %d\n", id)

	// 获取路由中的ID
	pid, err := strconv.ParseUint(c.Params("id", "0"), 10, 64)
	log.Printf("playlist id: %d\n", pid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "id error",
		})
	}

	// 获取查询参数
	type_ := c.Query("type", "")
	if type_ == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "not found query of type param",
		})
	}

	// 解析post json表单
	body := struct {
		SongId uint64 `json:"songId"`
	}{}
	// 解析失败，返回相应
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "json form error",
		})
	}
	log.Println(body)

	psIds := models.PlaylistSongIds{
		PlaylistId: pid,
	}

	switch type_ {
	// 添加歌曲到歌单
	case "add":
		if err := psIds.AddSong(body.SongId); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": fiber.StatusInternalServerError,
				"msg":  "add song to playlist failed",
			})
		}
	case "del":
		if err := psIds.DelSong(body.SongId); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": fiber.StatusInternalServerError,
				"msg":  "delete song to playlist failed",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"msg":  "success",
	})
}
