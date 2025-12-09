package day9

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		points = append(points, Point{x, y})
	}

	var segments []Segment
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		from := Point{
			x: int(math.Min(float64(p1.x), float64(p2.x))),
			y: int(math.Min(float64(p1.y), float64(p2.y))),
		}
		to := Point{
			x: int(math.Max(float64(p1.x), float64(p2.x))),
			y: int(math.Max(float64(p1.y), float64(p2.y))),
		}
		segments = append(segments, Segment{from, to})
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := computeRectanglesArea(points[i], points[j])
			if area > maxArea && checkRectangle(points[i], points[j], segments) {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func checkRectangle(a, b Point, segments []Segment) bool {
	minX := int(math.Min(float64(a.x), float64(b.x)))
	maxX := int(math.Max(float64(a.x), float64(b.x)))
	minY := int(math.Min(float64(a.y), float64(b.y)))
	maxY := int(math.Max(float64(a.y), float64(b.y)))

	for _, seg := range segments {
		from, to := seg.from, seg.to

		if from.y == to.y {
			y := from.y
			withinY := y > minY && y < maxY
			crossLeftEdge := minX >= from.x && minX < to.x
			crossRightEdge := maxX > from.x && maxX <= to.x

			if withinY && (crossLeftEdge || crossRightEdge) {
				return false
			}
		}

		if from.x == to.x {
			x := from.x
			withinX := x > minX && x < maxX
			crossTopEdge := minY >= from.y && minY < to.y
			crossBottomEdge := maxY > from.y && maxY <= to.y

			if withinX && (crossTopEdge || crossBottomEdge) {
				return false
			}
		}
	}

	return true
}
