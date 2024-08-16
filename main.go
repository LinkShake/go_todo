package main

import (
	"os"

	"github.com/LinkShake/go_todo/controllers"
	database "github.com/LinkShake/go_todo/db"
	"github.com/LinkShake/go_todo/helpers"
	"github.com/LinkShake/go_todo/middleware"
	"github.com/LinkShake/go_todo/templates/loginPage"
	"github.com/LinkShake/go_todo/templates/signupPage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Use(recover.New())

	app.Static("/public", "./public")

	app.Get("/is-user-logged-in", controllers.IsUserLoggedIn)
	app.Get("/", middleware.Auth, controllers.GetMainPage)
	app.Get("/_login", func (c *fiber.Ctx) error {
		if helpers.CheckLoggedIn(c) {
			return c.Redirect("/")
		}
		return helpers.Render(c, loginPage.LoginPage(""))
	})
	app.Get("/_signup", func (c *fiber.Ctx) error {
		if helpers.CheckLoggedIn(c) {
			return c.Redirect("/")
		}
		return helpers.Render(c, signupPage.SignUp(""))
	})
	app.Get("/get-edit-todo/:id", middleware.Auth, controllers.GetEditTodo)
	app.Post("/add-todo", middleware.Auth, controllers.AddTodo)
	app.Delete("/delete-todo", middleware.Auth, controllers.DeleteTodo)
	app.Put("/edit-todo", middleware.Auth, controllers.EditTodo)
	app.Get("/logout", middleware.Auth, controllers.Logout)
	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)

	app.Listen(":3000")	
}