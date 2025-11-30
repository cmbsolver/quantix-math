package main

import (
	"db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	db.InitDatabase()

	app := fiber.New()

	// Static files
	app.Static("/", "./quantix-math-ui/dist/quantix-math-ui/browser")

	app.Listen(":3000")
}
