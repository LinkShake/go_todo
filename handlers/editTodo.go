package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
)

func EditTodo(c *fiber.Ctx) error {
	db := database.DB
	body := new(schema.Todo)
	c.BodyParser(body)
	updatedTodo := &schema.Todo{ID: body.ID, UserId: body.UserId, Text: body.Text}
	res := db.Save(&updatedTodo)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.JSON(updatedTodo)
}