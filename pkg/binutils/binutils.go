package binutils

import (
	"fmt"
	"strconv"
	"strings"
)

func BytesToCSV(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	// Pre-size roughly: up to 4 chars per byte plus commas
	out := make([]byte, 0, len(b)*4)
	first := true
	for _, v := range b {
		if first {
			first = false
		} else {
			out = append(out, ',')
		}
		out = strconv.AppendInt(out, int64(v), 10)
	}
	return string(out)
}

func CSVToBytes(input string) ([]byte, error) {
	// Trim whitespace and newlines from the input
	input = strings.TrimSpace(input)

	// Split the input by commas
	valueStrings := strings.Split(input, ",")

	// Create a byte slice to hold the values
	bytes := make([]byte, 0, len(valueStrings))

	// Parse each value and add it to the byte slice
	for _, valueStr := range valueStrings {
		valueStr = strings.TrimSpace(valueStr)
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid integer\n", valueStr)
			return nil, err
		}

		if value < 0 || value > 255 {
			fmt.Printf("Error: Value %d is out of range (0-255)\n", value)
			return nil, fmt.Errorf("value out of range")
		}

		bytes = append(bytes, byte(value))
	}

	return bytes, nil
}
