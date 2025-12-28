package main

import (
	"quantix-math/pkg/db"
	"quantix-math/pkg/utility/loader"
	"quantix-math/routes" // Import the new routes package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// init database
	_, err := db.InitDatabase()
	if err != nil {
		return
	}
	loader.LoadWords()

	// Initialize standard Go html template engine
	engine := html.New("./views", ".tmpl")

	// Pass the engine to the Fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	// Setup Routes
	app.Static("/assets", "./assets")
	routes.RegisterAPIRoutes(app)
	routes.SetupUIRoutes(app)

	app.Listen(":3301")
}
