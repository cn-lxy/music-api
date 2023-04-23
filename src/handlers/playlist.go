package handlers

import (
	"log"

	"github.com/cn-lxy/music-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// PlaylistHandler is get user's playlist fiber handler
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
