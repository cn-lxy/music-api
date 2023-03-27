package handlers

import (
	"github.com/cn-lxy/music-api/tools"
	"github.com/gofiber/fiber/v2"
)

// VerfiyTokenHandler verify jwt token
func VerfiyTokenHandler(c *fiber.Ctx) error {
	// get token from header
	token := c.Get("Authorization")
	// verify token
	id, err := tools.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code": c.Status(fiber.StatusUnauthorized),
			"msg":  err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": c.Status(fiber.StatusOK),
		"msg":  "success",
		"data": id,
	})
}
