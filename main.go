package main

// @title Quantix Math API
// @version 1.0
// @description This is the Quantix Math API documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3301
// @BasePath /

import (
	"fmt"
	_ "quantix-math/docs"
	"quantix-math/pkg/db"
	"quantix-math/pkg/utility/loader"
	"quantix-math/routes" // Import the new routes package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // fiber-swagger middleware
	"github.com/gofiber/template/html/v2"
)

func main() {
	// init database
	_, err := db.InitDatabase()
	if err != nil {
		fmt.Printf("Database initialization failed: %v\n", err)
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
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	routes.RegisterAPIRoutes(app)
	routes.SetupUIRoutes(app)

	listenErr := app.Listen(":3301")
	if listenErr != nil {
		fmt.Printf("Server failed to start: %v\n", listenErr)
		return
	}
}
