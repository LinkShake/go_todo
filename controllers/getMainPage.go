package controllers

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/templates/mainPage"
	"github.com/LinkShake/go_todo/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []types.Todo
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

func GetMainPage(c *fiber.Ctx) error {
	var todos []schema.Todo
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
	return helpers.Render(c, mainPage.MainPage(todos))
}

func GetServerTodos(c *fiber.Ctx) ([]schema.Todo, error) {
	defer c.Response().CloseBodyStream()
	var todos []schema.Todo
	db := database.DB
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return []schema.Todo{}, err
	}
	res := db.Unscoped().Where("user_id = ?", parsedUserId).Order("created_at asc").Find(&todos)
	if res.Error != nil {
		return []schema.Todo{}, res.Error
	}
	return todos, nil
}