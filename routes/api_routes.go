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

	// Mobius routes
	apiGroup.Post("/mobius/direct", api.MobiusDirectHandler)
	apiGroup.Post("/mobius/mask", api.MobiusMaskHandler)
	apiGroup.Post("/mobius/divisor", api.MobiusDivisorHandler)

	// Runer routes
	apiGroup.Post("/runer/calculate-lines", api.CalculateLinesHandler)

	// Liber Primus routes
	apiGroup.Get("/liber-primus/metadata/:id", api.GetMetadataHandler)

	// Cipher analysis routes
	apiGroup.Post("/cipher/analyze", api.CipherAnalyzeHandler)
	apiGroup.Post("/cipher/matsui", api.MatsuiAnalyzeHandler)
}
