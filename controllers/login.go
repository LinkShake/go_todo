package controllers

import (
	"context"
	"os"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/redis"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/templates/loginPage"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

var UserCtx context.Context

func Login(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	db := database.DB
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	if email == "" || pwd == "" {
		return helpers.Render(c, loginPage.LoginPage("Received empty email or password"))
	}
	var user schema.User
	res := db.Unscoped().Where("email = ?", email).Find(&user)
	if res.RowsAffected == 0 {
		return helpers.Render(c, loginPage.LoginPage("Incorrect email or password"))
	}
	if res.Error != nil {
		if res.RowsAffected != 0 {
			panic(res.Error)
		}

		return helpers.Render(c, loginPage.LoginPage("Incorrect email or password"))
	}
	if match, err := argon2id.ComparePasswordAndHash(pwd, user.Pwd); err != nil {
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
			c.Response().Header.Set("HX-Redirect", "/")
			return c.SendStatus(fiber.StatusOK)
		}
		return helpers.Render(c, loginPage.LoginPage("Incorrect email or password"))
	}
}