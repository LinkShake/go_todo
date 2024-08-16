package controllers

import (
	"strconv"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/schema"
	"github.com/LinkShake/go_todo/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteTodo(c *fiber.Ctx) error {
	defer c.Response().CloseBodyStream()
	db := database.DB
	body := new(types.Todo)
	c.BodyParser(body)
	id := c.FormValue("todo-id")
	if id == "" {
		return c.SendString("not ok")
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
	todo := &schema.Todo{UserId: parsedUserId, ID: uint(parsedId)}
	res := db.Unscoped().Delete(&todo)
	if res.Error != nil {
		panic(res.Error)
	}
	return c.SendString(strconv.FormatUint(uint64(body.ID), 10))
}