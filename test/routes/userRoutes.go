package routes

import (
	"papernet/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	app.Get("/user", controllers.User)
}
