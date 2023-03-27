package middleware

import (
	"github.com/cn-lxy/music-api/tools/jwt"
	"github.com/gofiber/fiber/v2"
)

// VerifyMiddleware verify jwt token
func VerifyMiddleware(c *fiber.Ctx) error {
	path := c.Path()
	if path == "/login" || path == "/register" {
		return c.Next()
	}

	// get token from header
	token := c.Get("token")
	// verify token
	id, err := jwt.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code": c.Status(fiber.StatusUnauthorized),
			"msg":  err.Error(),
		})
	}
	// set user id to context
	c.Locals("id", id)
	return c.Next()
}
