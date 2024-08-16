package controllers

import (
	"strconv"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/templates/mainPage"
	"github.com/LinkShake/go_todo/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func EditTodo(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	db := database.DB
	body := new(types.Todo)
	c.BodyParser(body)
	id := c.FormValue("id")
	text := c.FormValue("cont")
	if text == "" && id == "" {
		return fiber.NewError(fiber.StatusNoContent, "Received empty input content and id")
	} else if text == "" && id != "" {
		return fiber.NewError(fiber.StatusNoContent, "Received empty input content")
	} else if text != "" && id == "" {
		return fiber.NewError(fiber.StatusNoContent, "Received empty id")
	}
	parsedId, parsingErr := strconv.ParseUint(id, 10, 64)
	if parsingErr != nil {
		panic(parsingErr)
	}
	userId := c.Locals("userId").(string)
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(err)
	}
	updatedTodo := &schema.Todo{ID: uint(parsedId), UserId: parsedUserId, Text: text}
	res := db.Save(&updatedTodo)
	if res.Error != nil {
		panic(res.Error)
	}
	return helpers.Render(c, mainPage.Todo(*updatedTodo, false))
}