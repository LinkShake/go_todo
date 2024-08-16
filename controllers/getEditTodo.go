package controllers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/templates/mainPage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetEditTodo(c *fiber.Ctx) error {
	db := database.DB
	todoId := c.Params("id")
	if todoId == "" {
		return fiber.NewError(fiber.StatusNoContent, "Received empty id")
	}
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	var todo schema.Todo
	res := db.Unscoped().Where("user_id = ?", parsedUserId).Where("id = ?", todoId).Find(&todo)
	if res.Error != nil {
		panic(res.Error)
	}
	return helpers.Render(c, mainPage.Todo(todo, true))
}