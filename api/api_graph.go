package api

import (
	"github.com/gofiber/fiber/v2"
	"math"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Line struct {
	P1     Point   `json:"p1"`
	P2     Point   `json:"p2"`
	Length float64 `json:"length"`
}

type GraphRequest struct {
	Points []Point `json:"points"`
}

type GraphResponse struct {
	Points      []Point `json:"points"`
	Lines       []Line  `json:"lines"`
	TotalLength float64 `json:"totalLength"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// ProcessGraphHandler handles the graph processing request
// @Summary Process a set of coordinate points
// @Description Calculates lines between points and total length, filtering lines that pass through other points
// @Tags Math
// @Accept  json
// @Produce  json
// @Param   request  body      GraphRequest  true  "Graph Request"
// @Success 200      {object}  GraphResponse
// @Failure 400      {object}  ErrorResponse
// @Router /api/graph/process [post]
func ProcessGraphHandler(c *fiber.Ctx) error {
	var req GraphRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Cannot parse request body",
		})
	}

	points := req.Points
	if len(points) < 2 {
		return c.JSON(GraphResponse{
			Points:      points,
			Lines:       []Line{},
			TotalLength: 0,
		})
	}

	var lines []Line
	totalLength := 0.0

	// Generate all possible lines and filter them
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			// Check if this line passes through any other point
			passesThrough := false
			for k := 0; k < len(points); k++ {
				if k == i || k == j {
					continue
				}
				if pointOnLine(points[k], p1, p2) {
					passesThrough = true
					break
				}
			}

			if !passesThrough {
				length := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))

				// Ensure consistent ordering of points in Line to avoid "repeating sets"
				// though our loops already avoid duplicates by having j > i.
				// But just in case points themselves are duplicates or swapped.
				orderedP1, orderedP2 := orderPoints(p1, p2)

				lines = append(lines, Line{
					P1:     orderedP1,
					P2:     orderedP2,
					Length: length,
				})
				totalLength += length
			}
		}
	}

	return c.JSON(GraphResponse{
		Points:      points,
		Lines:       lines,
		TotalLength: totalLength,
	})
}

// pointOnLine checks if point p lies on the line segment between a and b
func pointOnLine(p, a, b Point) bool {
	// Check if p is between a and b (bounding box check)
	if p.X < math.Min(a.X, b.X) || p.X > math.Max(a.X, b.X) ||
		p.Y < math.Min(a.Y, b.Y) || p.Y > math.Max(a.Y, b.Y) {
		return false
	}

	// Cross product to check collinearity (with some tolerance for float precision)
	// (y2-y1)(x3-x1) == (y3-y1)(x2-x1)
	crossProduct := (p.Y-a.Y)*(b.X-a.X) - (p.X-a.X)*(b.Y-a.Y)

	const epsilon = 1e-9
	return math.Abs(crossProduct) < epsilon
}

func orderPoints(p1, p2 Point) (Point, Point) {
	if p1.X < p2.X || (p1.X == p2.X && p1.Y < p2.Y) {
		return p1, p2
	}
	return p2, p1
}
