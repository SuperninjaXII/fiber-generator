package main

import (
	"{AppName}/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "fiber app",
		})
	})
	routes.User(app)

	log.Fatal(app.Listen(":3000"))
}
