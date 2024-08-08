package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
)

func AddTodo(c *fiber.Ctx) error {
	db := database.DB
	body := new(schema.Todo)
	c.BodyParser(body)
	newTodo := &schema.Todo{Text: body.Text, UserId: body.UserId}
	res := db.Create(&newTodo)
	if res.Error != nil {
		panic(res.Error)
	}
	return	c.JSON(newTodo)
}