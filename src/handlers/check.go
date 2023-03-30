package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// CheckdHandler
func CheckdHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["id"]
	name := claims["nickname"].(string)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "OK ",
		"code": fiber.StatusOK,
		"data": fiber.Map{
			"name": name,
			"id":   id,
		},
	})
}
