package main

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.DB
	database.Migrate(db)

	app := fiber.New()

	app.Get("/todos/:userId", handlers.GetTodos)
	app.Post("/add-todo", handlers.AddTodo)
	app.Delete("/delete-todo", handlers.DeleteTodo)
	app.Put("/edit-todo", handlers.EditTodo)

	app.Listen(":3000")	
}