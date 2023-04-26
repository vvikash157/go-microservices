package geometry

import (
	"math"
)

func Area(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length * width)
	return
}

func Diagonal(length, width float64) (diagonal float64) {
	diagonal = math.Sqrt((length * length) + (width * width))
	return
}
