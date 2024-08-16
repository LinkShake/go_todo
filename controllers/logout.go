package controllers

import (
	"github.com/LinkShake/go_todo/redis"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	sid := c.Cookies("sid")
	err := redis.RemoveSessionId(sid)
	if err != nil {
		panic(err)
	}
	c.ClearCookie("sid")
	c.Locals("userId", "")
	c.Response().Header.Set("HX-Redirect", "/_login")
	return c.SendStatus(fiber.StatusOK)
}