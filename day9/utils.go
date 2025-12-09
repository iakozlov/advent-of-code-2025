package day9

import "math"

type Point struct {
	x, y int
}

type Segment struct {
	from, to Point
}

func computeRectanglesArea(p1 Point, p2 Point) int {
	xLen := math.Abs(float64(p1.x-p2.x)) + 1
	yLen := math.Abs(float64(p1.y-p2.y)) + 1
	return int(xLen * yLen)
}
