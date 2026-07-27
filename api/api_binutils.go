package api

import (
	"io"
	"quantix-math/pkg/binutils"

	"github.com/gofiber/fiber/v2"
)

type BytesToCSVResponse struct {
	CSV string `json:"csv"`
}

type CSVToBytesRequest struct {
	CSV string `json:"csv"`
}

type Base64ToCSVRequest struct {
	Base64 string `json:"base64"`
}

type CSVToBase64Request struct {
	CSV string `json:"csv"`
}

type CSVToBase64Response struct {
	Base64 string `json:"base64"`
}

// FileToCSVHandler handles the file to CSV conversion request
// @Summary Convert a file to hex CSV
// @Description Reads a file and returns its content as a hex-encoded CSV string
// @Tags Utilities
// @Accept  multipart/form-data
// @Produce  json
// @Param   file  formData  file  true  "File to convert"
// @Success 200   {object}  BytesToCSVResponse
// @Failure 400   {string}  string "File upload required"
// @Failure 500   {string}  string "Internal server error"
// @Router /api/binutils/file-to-csv [post]
func FileToCSVHandler(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("File upload required")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to open file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read file content")
	}

	csv := binutils.BytesToCSV(content)
	return c.JSON(BytesToCSVResponse{CSV: csv})
}

// CSVToBytesHandler handles the CSV to binary conversion request
// @Summary Convert hex CSV to binary
// @Description Converts a hex-encoded CSV string back to binary data and returns it as a file
// @Tags Utilities
// @Accept  json
// @Produce  application/octet-stream
// @Param   request  body      CSVToBytesRequest  true  "CSV to Bytes Request"
// @Success 200      {file}    binary
// @Failure 400      {string}  string "Invalid request body or CSV format"
// @Router /api/binutils/csv-to-bytes [post]
func CSVToBytesHandler(c *fiber.Ctx) error {
	var req CSVToBytesRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	bytes, err := binutils.CSVToBytes(req.CSV)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Set appropriate headers for file download
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename=\"output.bin\"")

	return c.Send(bytes)
}

// Base64ToCSVHandler handles the Base64 to CSV conversion request
// @Summary Convert Base64 to hex CSV
// @Description Converts a Base64 string to a hex-encoded CSV string
// @Tags Utilities
// @Accept  json
// @Produce  json
// @Param   request  body      Base64ToCSVRequest  true  "Base64 to CSV Request"
// @Success 200      {object}  BytesToCSVResponse
// @Failure 400      {string}  string "Invalid request body or Base64 format"
// @Router /api/binutils/base64-to-csv [post]
func Base64ToCSVHandler(c *fiber.Ctx) error {
	var req Base64ToCSVRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	bytes, err := binutils.Base64ToBytes(req.Base64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	csv := binutils.BytesToCSV(bytes)
	return c.JSON(BytesToCSVResponse{CSV: csv})
}

// CSVToBase64Handler handles the CSV to Base64 conversion request
// @Summary Convert hex CSV to Base64
// @Description Converts a hex-encoded CSV string to a Base64 string
// @Tags Utilities
// @Accept  json
// @Produce  json
// @Param   request  body      CSVToBase64Request  true  "CSV to Base64 Request"
// @Success 200      {object}  CSVToBase64Response
// @Failure 400      {string}  string "Invalid request body or CSV format"
// @Router /api/binutils/csv-to-base64 [post]
func CSVToBase64Handler(c *fiber.Ctx) error {
	var req CSVToBase64Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	bytes, err := binutils.CSVToBytes(req.CSV)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	b64 := binutils.BytesToBase64(bytes)
	return c.JSON(CSVToBase64Response{Base64: b64})
}
