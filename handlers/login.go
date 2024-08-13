package handlers

import (
	"context"
	"os"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/redis"
	"github.com/LinkShake/go_todo/schema"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

var UserCtx context.Context

func Login(c *fiber.Ctx) error {
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	db := database.DB
	body := new(User)
	c.BodyParser(body)
	var user schema.User
	res := db.Unscoped().Where("email = ?", body.Email).Find(&user)
	if res.RowsAffected == 0 {
		return c.JSON(&ReqFailed{
			Ok: false,
			Msg: "user not found",
		})
	}
	if res.Error != nil {
		if res.RowsAffected != 0 {
			panic(res.Error)
		}

		return c.JSON(&ReqFailed{
			Ok: false,
			Msg: "user not found",
		})
	}
	if match, err := argon2id.ComparePasswordAndHash(body.Pwd, user.Pwd); err != nil {
		panic(err)
	} else {
		if match {
			sid, _, redisErr := redis.UpdateUserSessionIdx(user.ID.String())
			if redisErr != nil {
				panic(redisErr)
			}
			c.Cookie(&fiber.Cookie{
				Name: "sid",
				Value: sid,
				HTTPOnly: true,
				Secure: os.Getenv("ENV") == "production",
				SameSite: "lax",
				Path: "/",
				Domain: getDomain(),
				MaxAge: 1000 * 60 * 60 * 24 * 365 * 10,
			})
			return c.Redirect("/")
		}
		return c.SendString("invalid pwd")
	}
}