package handlers

import "github.com/gofiber/fiber/v2"

// PlaylistHandler is get user's playlist fiber handler
// the api is "/playlist"
func PlaylistHandler(c *fiber.Ctx) error {
	// get user's playlist
	// playlist, err := GetPlaylist()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Error getting playlist",
	// 	})
	// }

	// return user's playlist
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		// "data":    playlist,
	})
}
