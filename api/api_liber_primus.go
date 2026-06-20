package api

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

// ImageMetadata represents the metadata for an image
type ImageMetadata struct {
	Description    string         `json:"description"`
	Dimensions     string         `json:"dimensions"`
	Color          string         `json:"color"`
	ColorCounts    map[string]int `json:"color_counts"`
	ICCProfile     string         `json:"icc_profile"`
	ICCDescription string         `json:"icc_description,omitempty"`
	JFIF           *JFIFMetadata  `json:"jfif,omitempty"`
}

type JFIFMetadata struct {
	Version  string `json:"version"`
	Units    string `json:"units"`
	XDensity int    `json:"x_density"`
	YDensity int    `json:"y_density"`
}

func analyzeImage(img image.Image) (string, map[string]int) {
	bounds := img.Bounds()

	colorCounts := make(map[string]int)
	isColor := false

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Use 8-bit for comparison
			r8 := r >> 8
			g8 := g >> 8
			b8 := b >> 8

			// Format color as hex
			colorHex := fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)
			colorCounts[colorHex]++

			if abs(int(r8)-int(g8)) > 5 || abs(int(r8)-int(b8)) > 5 || abs(int(g8)-int(b8)) > 5 {
				isColor = true
			}
		}
	}

	colorType := "Grayscale"
	if isColor {
		colorType = "Color"
	}

	return colorType, colorCounts
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func extractJFIF(data []byte) (*JFIFMetadata, error) {
	// Find APP0 marker (FF E0)
	idx := bytes.Index(data, []byte{0xFF, 0xE0})
	if idx == -1 || idx+16 >= len(data) {
		return nil, fmt.Errorf("JFIF marker not found")
	}

	// Check for "JFIF\0" identifier
	if !bytes.Equal(data[idx+4:idx+9], []byte("JFIF\x00")) {
		return nil, fmt.Errorf("not a JFIF header")
	}

	// Parse
	major := data[idx+9]
	minor := data[idx+10]
	units := data[idx+11]
	xDensity := int(data[idx+12])<<8 | int(data[idx+13])
	yDensity := int(data[idx+14])<<8 | int(data[idx+15])

	unitStr := "None"
	switch units {
	case 1:
		unitStr = "dots/inch"
	case 2:
		unitStr = "dots/cm"
	}

	return &JFIFMetadata{
		Version:  fmt.Sprintf("%d.%d", major, minor),
		Units:    unitStr,
		XDensity: xDensity,
		YDensity: yDensity,
	}, nil
}

func GetMetadataHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	imagePath := fmt.Sprintf("assets/files/images/%s.jpg", id)
	file, err := os.Open(imagePath)
	if err != nil {
		return c.JSON(ImageMetadata{
			Description: fmt.Sprintf("Liber Primus - Page %s", id),
			Dimensions:  "N/A",
			Color:       "N/A",
			ICCProfile:  "N/A",
		})
	}
	defer func() { _ = file.Close() }()

	data, err := io.ReadAll(file)
	if err != nil {
		return c.JSON(ImageMetadata{
			Description: fmt.Sprintf("Liber Primus - Page %s", id),
			Dimensions:  "N/A",
			Color:       "N/A",
			ICCProfile:  "N/A",
		})
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return c.JSON(ImageMetadata{
			Description: fmt.Sprintf("Liber Primus - Page %s", id),
			Dimensions:  "N/A",
			Color:       "N/A",
			ICCProfile:  "N/A",
		})
	}

	bounds := img.Bounds()
	dimensions := fmt.Sprintf("%dx%d", bounds.Dx(), bounds.Dy())
	color, colorCounts := analyzeImage(img)

	iccProfile := "None"
	iccDescription := ""
	if bytes.Contains(data, []byte("ICC_PROFILE")) {
		iccProfile = "Embedded"
		// Simple attempt to extract description
		if idx := bytes.Index(data, []byte("Artifex Software")); idx != -1 {
			// Find end of the string (null byte)
			end := bytes.Index(data[idx:], []byte{0})
			if end != -1 {
				iccDescription = string(data[idx : idx+end])
			} else {
				iccDescription = "Artifex Software sRGB ICC Profile"
			}
		}
	}

	metadata := ImageMetadata{
		Description:    fmt.Sprintf("Liber Primus - Page %s", id),
		Dimensions:     dimensions,
		Color:          color,
		ColorCounts:    colorCounts,
		ICCProfile:     iccProfile,
		ICCDescription: iccDescription,
	}

	if jfif, err := extractJFIF(data); err == nil {
		metadata.JFIF = jfif
	}

	return c.JSON(metadata)
}
