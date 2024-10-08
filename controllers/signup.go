package controllers

import (
	"fmt"
	"os"
	"strings"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/redis"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/templates/signupPage"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Signup(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	db := database.DB
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	if email == "" || pwd == "" {
		return helpers.Render(c, signupPage.SignUp("Received empty email or password"))
	} 
	if !strings.Contains(email, "@") {
		return helpers.Render(c, signupPage.SignUp("Invalid email format"))
	}
	var oldUser schema.User
	res := db.Unscoped().Where("email = ?", email).First(&oldUser)
	if res.Error != nil {
		if res.RowsAffected != 0 {
			return helpers.Render(c, signupPage.SignUp("Email already in use"))
		}

		hash, err := argon2id.CreateHash(pwd, argon2id.DefaultParams)
		if err != nil {
			panic(err)
		}
		newUser := &schema.User{Email: email, Pwd: hash}
		res := db.Create(&newUser)
		if res.Error != nil {
			panic(res.Error)
		}
		sid, _, redisErr := redis.UpdateUserSessionIdx(newUser.ID.String())
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
	return c.SendStatus(fiber.StatusOK)
}

func getDomain() string {
	if os.Getenv("ENV") == "production" {
		return fmt.Sprintf(".${%v}", os.Getenv("DOMAIN"))
	} 
	return ""
}