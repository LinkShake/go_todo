package handlers

import (
	"fmt"
	"os"
	"strings"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/redis"
	"github.com/LinkShake/go_todo/schema"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Signup(c *fiber.Ctx) error {
	isUserLoggedIn := helpers.CheckLoggedIn(c)
	if isUserLoggedIn {
		return c.Redirect("/")
	}
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	db := database.DB
	body := new(User)
	c.BodyParser(body)
	if !strings.Contains(body.Email, "@") {
		return c.JSON(&ReqFailed{
			Ok: false,
			Msg: "invalid email",
		})
	}
	var oldUser schema.User
	res := db.Unscoped().Where("email = ?", body.Email).First(&oldUser)
	if res.Error != nil {
		if res.RowsAffected != 0 {
			panic(res.Error)
		}

		hash, err := argon2id.CreateHash(body.Pwd, argon2id.DefaultParams)
		if err != nil {
			panic(err)
		}
		newUser := &schema.User{Email: body.Email, Pwd: hash}
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
		return c.Redirect("/")
	}
	return c.SendString("")
}

func getDomain() string {
	if os.Getenv("ENV") == "production" {
		return fmt.Sprintf(".${%v}", os.Getenv("DOMAIN"))
	} 
	return ""
}