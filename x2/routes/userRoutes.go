package routes

import (
	"x2/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	app.Get("/user", controllers.User)
}
