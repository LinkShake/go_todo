package handlers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []schema.Todo
	db := database.DB
	userId := c.Params("userId")
	res := db.Unscoped().Where("user_id = ?", userId).Order("created_at asc").Find(&todos)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.JSON(todos)
}