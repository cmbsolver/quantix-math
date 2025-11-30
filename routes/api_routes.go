package routes

import (
	"quantix-math/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/sequence", api.GetSequenceHandler)
}
