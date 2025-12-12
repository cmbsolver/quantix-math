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

	app.Get("/isitprime", func(c *fiber.Ctx) error {
		return c.Render("isitprime", fiber.Map{
			"Title": "Is It Prime?",
		})
	})

	app.Get("/file_to_csv", func(c *fiber.Ctx) error {
		return c.Render("file_to_csv", fiber.Map{
			"Title": "File to CSV",
		})
	})

	app.Get("/csv_to_bytes", func(c *fiber.Ctx) error {
		return c.Render("csv_to_bytes", fiber.Map{
			"Title": "CSV to Bytes",
		})
	})

	app.Get("/runecalc", func(c *fiber.Ctx) error {
		return c.Render("runecalc", fiber.Map{
			"Title": "Rune Calculator",
		})
	})

	app.Get("/dictionary_words", func(c *fiber.Ctx) error {
		return c.Render("dictionary_words", fiber.Map{
			"Title": "Dictionary Words",
		})
	})
}
