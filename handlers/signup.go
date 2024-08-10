package handlers

import (
	"strings"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	db := database.DB
	body := new(User)
	c.BodyParser(body)
	if !strings.Contains(body.Email, "@") {
		panic("Invalid email")
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
		return c.JSON(newUser)
	}
	return c.SendString("")
}