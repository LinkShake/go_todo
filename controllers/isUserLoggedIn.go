package controllers

import (
	"github.com/LinkShake/go_todo/helpers"
	"github.com/gofiber/fiber/v2"
)

func IsUserLoggedIn(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	return c.SendStatus(fiber.StatusOK)
}