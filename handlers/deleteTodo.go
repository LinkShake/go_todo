package handlers

import (
	"strconv"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB
	body := new(Todo)
	c.BodyParser(body)
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	todo := &schema.Todo{UserId: parsedUserId, ID: body.ID}
	res := db.Unscoped().Delete(&todo)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.SendString(strconv.FormatUint(uint64(body.ID), 10))
}