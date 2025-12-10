package day10

import (
	"bufio"
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

	totalPresses := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		targetMask, buttons, err := parseLinePart2(line)
		if err != nil {
			return 0, err
		}
		presses := minPressesJoltageBFS(targetMask, buttons)
		totalPresses += presses
	}

	return totalPresses, nil
}

func encodeState(slice []int) string {
	if len(slice) == 0 {
		return ""
	}
	var b strings.Builder
	b.Grow(len(slice) * 4)
	for i, v := range slice {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

func minPressesJoltageBFS(targets []int, buttons [][]int) int {
	start := make([]int, len(targets))
	if equalSlice(start, targets) {
		return 0
	}

	visited := make(map[string]struct{})
	startKey := encodeState(start)
	visited[startKey] = struct{}{}

	var queue [][]int
	queue = append(queue, start)

	presses := 0
	for len(queue) > 0 {
		var nextQueue [][]int

		for _, state := range queue {
			for _, btn := range buttons {
				ok := true
				for _, idx := range btn {
					if state[idx] >= targets[idx] {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}
				newState := make([]int, len(targets))
				copy(newState, state)
				for _, idx := range btn {
					newState[idx]++
				}

				if equalSlice(newState, targets) {
					return presses + 1
				}

				key := encodeState(newState)
				if _, seen := visited[key]; !seen {
					visited[key] = struct{}{}
					nextQueue = append(nextQueue, newState)
				}
			}
		}
		queue = nextQueue
		presses++
	}

	return -1
}
