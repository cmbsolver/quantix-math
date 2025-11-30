package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupUIRoutes initializes the views/UI related routes
func SetupUIRoutes(app *fiber.App) {
	// Home Page
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home",
		})
	})

	// About Page
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "About",
		})
	})

	// About Page
	app.Get("/sequence", func(c *fiber.Ctx) error {
		return c.Render("sequence", fiber.Map{
			"Title": "Numeric Sequences",
		})
	})
}
