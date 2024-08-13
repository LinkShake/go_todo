package handlers

import (
	"github.com/LinkShake/go_todo/helpers"
	"github.com/gofiber/fiber/v2"
)

func IsUserLoggedIn(c *fiber.Ctx) error {
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	return c.SendString("ok")
}