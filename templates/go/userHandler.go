package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func User(c *fiber.Ctx) error {
	return c.Render("user", fiber.Map{
		"Text": "users page",
	})
}
