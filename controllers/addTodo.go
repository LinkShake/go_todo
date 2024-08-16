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

func AddTodo(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	db := database.DB
	body := new(types.Todo)
	c.BodyParser(body)
	text := c.FormValue("cont")
	if text == "" {
		return fiber.NewError(fiber.StatusNoContent, "Received empty input content")
	}
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	newTodo := &schema.Todo{Text: text, UserId: parsedUserId}
	res := db.Create(&newTodo)
	if res.Error != nil {
		panic(res.Error)
	}
	return	helpers.Render(c, mainPage.Todo(*newTodo, false))
}