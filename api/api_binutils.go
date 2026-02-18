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
