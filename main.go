package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Static files
	app.Static("/", "./quantix-math-ui/dist/quantix-math-ui/browser")

	app.Listen(":3000")
}
