package main

import (
	"log"
	"{AppName}/routes"

	"github.com/aarol/reload"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	isDevelopment := true

	// Initialize the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files from the "public" folder
	app.Static("/", "./public")

	// Define a route to render the index page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Fiber App with Reload",
		})
	})

	// Setup user-defined routes
	routes.User(app)

	// Add the reload middleware in development mode
	if isDevelopment {
		reloader := reload.New("views/", "public/") // Watch views and public directories
		reloader.OnReload = func() {
			log.Println("Reload triggered")
		}

		reloadMiddleware := fiber.WrapHTTPHandler(reloader.Handle(app.Handler()))
		app.Use(func(c *fiber.Ctx) error {
			reloadMiddleware.ServeHTTP(c.Context(), c.Context().Request())
			return nil
		})
	}

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

