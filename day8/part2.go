package day8

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []Point
	for scanner.Scan() {
		line := scanner.Text()
		coordinates := strings.Split(line, ",")
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return 0, err
		}
		z, err := strconv.Atoi(coordinates[2])
		if err != nil {
			return 0, err
		}
		points = append(points, Point{x, y, z})
	}
	edges := buildEdges(points)
	sort.Slice(edges, func(i, j int) bool { return edges[i].squaredDistance < edges[j].squaredDistance })

	d := NewDSU(len(points))
	components := len(points)
	var lastMerdgedEdge Edge

	for _, edge := range edges {
		isMerged := d.Union(edge.i, edge.j)
		if isMerged {
			components--
		}
		if components == 1 {
			lastMerdgedEdge = edge
			break
		}
	}

	x1 := points[lastMerdgedEdge.i].x
	x2 := points[lastMerdgedEdge.j].x
	return x1 * x2, nil
}
