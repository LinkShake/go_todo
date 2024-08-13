package main

import (
	"os"

	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/handlers"
	"github.com/LinkShake/go_todo/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/joho/godotenv"
)

func main() {
	db := database.DB
	database.Migrate(db)

	app := fiber.New()
	
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}

	app.Use(encryptcookie.New(encryptcookie.Config{
    	Key: os.Getenv("COOKIE_ENCRYPTION_KEY"),
	}))

	app.Static("/_login", "./public/_login", fiber.Static{
		CacheDuration: -1,
	})
	app.Static("/_signup", "./public/_signup", fiber.Static{
		CacheDuration: -1,
	})
	app.Static("/", "./public", fiber.Static{
 	   CacheDuration: -1,
	})

	app.Get("/is-user-logged-in", handlers.IsUserLoggedIn)
	app.Get("/todos", middleware.Auth, handlers.GetTodos)
	app.Post("/add-todo", middleware.Auth, handlers.AddTodo)
	app.Delete("/delete-todo", middleware.Auth, handlers.DeleteTodo)
	app.Put("/edit-todo", middleware.Auth, handlers.EditTodo)
	app.Get("/logout", middleware.Auth, handlers.Logout)
	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)

	app.Listen(":3000")	
}