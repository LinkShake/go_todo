package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB
	body := new(schema.Todo)
	c.BodyParser(body)
	todo := &schema.Todo{UserId: body.UserId, ID: body.ID}
	res := db.Delete(&todo)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.SendString("ok")
}