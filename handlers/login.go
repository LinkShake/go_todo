package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email string
	Pwd string
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	body := new(User)
	c.BodyParser(body)
	var user schema.User
	res := db.Unscoped().Where("email = ?", body.Email).Find(&user)
	if res.RowsAffected == 0 {
		return c.SendString("user not found")
	}
	if res.Error != nil {
		if res.RowsAffected != 0 {
			panic(res.Error)
		}

		return c.SendString("user not found")
	}
	if match, err := argon2id.ComparePasswordAndHash(body.Pwd, user.Pwd); err != nil {
		panic(err)
	} else {
		if match {
			return c.JSON(user)
		}
		return c.SendString("invalid pwd")
	}
}