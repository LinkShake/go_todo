package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func EditTodo(c *fiber.Ctx) error {
	db := database.DB
	body := new(Todo)
	c.BodyParser(body)
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	updatedTodo := &schema.Todo{ID: body.ID, UserId: parsedUserId, Text: body.Text}
	res := db.Save(&updatedTodo)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.JSON(updatedTodo)
}