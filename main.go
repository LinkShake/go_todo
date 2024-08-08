package main

import (
	database "github.com/LinkShake/go_todo/db"
	"github.com/gofiber/fiber"
)

func main() {
	db := database.Connect()
	database.Migrate(db)

	app := fiber.New()

	app.Listen(3000)	
}