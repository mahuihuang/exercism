package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case r <= 1:
		return 10
	case r <= 5:
		return 5
	case r <= 10:
		return 1
	default:
		return 0
	}
}
