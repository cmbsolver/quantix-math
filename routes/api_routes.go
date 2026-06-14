package routes

import (
	"quantix-math/api"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/sequence", api.GetSequenceHandler)
	apiGroup.Post("/check-number", api.CheckNumberInSequencesHandler)
	apiGroup.Post("/prime", api.GetIsItPrimeHandler)

	// Dictionary routes
	apiGroup.Post("/dictionary/words", api.GetDictionaryWordsByParamHandler)
	apiGroup.Get("/dictionary/words/download", api.DownloadDictionaryWordsExcelHandler)
	apiGroup.Post("/dictionary/anagrams", api.GetAnagramsHandler)

	// Binutils routes
	apiGroup.Post("/binutils/file-to-csv", api.FileToCSVHandler)
	apiGroup.Post("/binutils/csv-to-bytes", api.CSVToBytesHandler)
	apiGroup.Post("/binutils/base64-to-csv", api.Base64ToCSVHandler)
	apiGroup.Post("/binutils/csv-to-base64", api.CSVToBase64Handler)

	// Graph routes
	apiGroup.Post("/graph/process", api.ProcessGraphHandler)

	// Runer routes
	apiGroup.Post("/runer/calculate-lines", api.CalculateLinesHandler)
}
