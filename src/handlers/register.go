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
	if err := user.Insert(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"msg":  err.Error(),
		})
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
