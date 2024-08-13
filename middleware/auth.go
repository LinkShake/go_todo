package middleware

import (
	"github.com/LinkShake/go_todo/redis"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	sid := c.Cookies("sid")
	if sid == "" {
		return c.Redirect("/_login")
	}
	userId, err := redis.GetUserId(sid)
	if err != nil {
		if err.Error() == "invalid session id" {
			return c.Redirect("/_login")
		}
		panic(err)
	}
	c.Locals("userId", userId)
	return c.Next()
}
