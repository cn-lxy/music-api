package handlers

import (
	"github.com/cn-lxy/music-api/models"
	"github.com/cn-lxy/music-api/tools"
	"github.com/gofiber/fiber/v2"
)

// email and password struct
type formEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// nickname and password struct
type formNickName struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

// LoginHandler User login handler
func LoginHandler(c *fiber.Ctx) error {
	var form any
	t := c.Query("t")
	if t == "nickname" {
		form = &formNickName{}
	} else if t == "email" {
		form = &formEmail{}
	} else if t == "" {
		// the query params of `t` is required
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": c.Status(fiber.StatusBadRequest),
			"msg":  "the query params of `t` is required",
		})
	} else {
		// the query params of `t` is `email` or `nickname`.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": c.Status(fiber.StatusBadRequest),
			"msg":  "the query params of `t` is valid, the value should `email` or `nickname`",
		})
	}
	// create user
	user := &models.User{}
	// parse body into user model
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": c.Status(fiber.StatusBadRequest),
			"msg":  err.Error(),
		})
	}
	if t == "nickname" {
		if f, ok := form.(*formNickName); ok {
			user.NickName = f.Nickname
			user.Password = f.Password
		}
	}
	if t == "email" {
		if f, ok := form.(*formEmail); ok {
			user.Email = f.Email
			user.Password = f.Password
		}
	}
	// get user from database
	if err := user.GetByEmailOrNick(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code": c.Status(fiber.StatusUnauthorized),
			"msg":  err.Error(),
		})
	}
	// generate JWT token
	token, err := tools.GenerateToken(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": c.Status(fiber.StatusInternalServerError),
			"msg":  err.Error(),
		})
	}

	// return JWT token and code
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": c.Status(fiber.StatusOK),
		"msg":  "success",
		"data": fiber.Map{
			"token": token,
		},
	})
}
