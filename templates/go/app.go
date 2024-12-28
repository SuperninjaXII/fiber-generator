package main

import (
	"fmt"
	"log"
	"net/"
	"strings"
	"{AppName}/routes"

	"github.com/aarol/reload"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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

	if isDevelopment {
		// Create a new reloader instance
		reloader := reload.New("views/", "public/")

		// Optional: Log when reload is triggered
		reloader.OnReload = func() {
			log.Println("Reload triggered")
		}

		// Create WebSocket endpoint for reload notifications
		app.Get("/reload_ws", adaptor.HTTPHandler(
			http.HandlerFunc(reloader.ServeWS),
		))

		// Add middleware to inject reload script
		app.Use(func(c *fiber.Ctx) error {
			// Only inject script into HTML responses
			if strings.Contains(c.Get("Content-Type"), "text/html") {
				response := c.Response()
				body := response.Body()

				// Inject the reload script before closing </body> tag
				script := reload.InjectedScript("/reload_ws")
				newBody := strings.Replace(
					string(body),
					"</body>",
					fmt.Sprintf("%s</body>", script),
					1,
				)

				response.SetBody([]byte(newBody))
			}

			// Disable caching in development
			c.Set("Cache-Control", "no-cache")

			return c.Next()
		})

		// Start the file watcher
		go reloader.WatchDirectories()
	}

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Fiber App with Reload",
		})
	})

	// Setup user-defined routes
	routes.User(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

