package handlers

import (
	"github.com/LinkShake/go_todo/redis"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	sid := c.Cookies("sid")
	err := redis.RemoveSessionId(sid)
	if err != nil {
		return c.JSON(&ReqFailed{
			Ok: false,
			Msg: err.Error(),
		})
	}
	c.ClearCookie("sid")
	c.Locals("userId", "")
	return c.Redirect("/_login")
}