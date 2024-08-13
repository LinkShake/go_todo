package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo
	db := database.DB
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	res := db.Unscoped().Where("user_id = ?", parsedUserId).Order("created_at asc").Find(&todos)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.JSON(todos)
}