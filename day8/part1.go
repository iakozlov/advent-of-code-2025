package day8

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
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

	const Runs = 1000
	d := NewDSU(len(points))
	for i := 0; i < Runs; i++ {
		e := edges[i]
		d.Union(e.i, e.j)
	}

	componentsMap := make(map[int]int)
	for i := 0; i < len(points); i++ {
		componentsMap[d.Find(i)]++
	}

	sizes := make([]int, 0, len(componentsMap))
	for _, size := range componentsMap {
		sizes = append(sizes, size)
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	a := sizes[0]
	b := 1
	c := 1
	if len(sizes) > 1 {
		b = sizes[1]
	}
	if len(sizes) > 2 {
		c = sizes[2]
	}

	return a * b * c, nil
}
