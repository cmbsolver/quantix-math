package routes

import (
	"quantix-math/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/sequence", api.GetSequenceHandler)
	apiGroup.Post("/prime", api.GetIsItPrimeHandler)

	// Dictionary routes
	apiGroup.Post("/dictionary/words", api.GetDictionaryWordsByParamHandler)

	// Binutils routes
	apiGroup.Post("/binutils/file-to-csv", api.FileToCSVHandler)
	apiGroup.Post("/binutils/csv-to-bytes", api.CSVToBytesHandler)
}
