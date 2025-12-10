package day10

import (
	"bufio"
	"os"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalPresses := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		targetMask, buttons, err := parseLinePart1(line)
		if err != nil {
			return 0, err
		}
		presses := minPressesBFS(targetMask, buttons)
		totalPresses += presses
	}

	return totalPresses, nil
}

func minPressesBFS(target uint64, buttons []uint64) int {
	visited := make(map[uint64]int)
	var queue []uint64

	startState := uint64(0)
	visited[startState] = 0
	queue = append(queue, startState)

	for head := 0; head < len(queue); head++ {
		state := queue[head]
		dist := visited[state]
		nextDist := dist + 1

		for _, button := range buttons {
			ns := state ^ button
			if _, seen := visited[ns]; seen {
				continue
			}
			if ns == target {
				return nextDist
			}
			visited[ns] = nextDist
			queue = append(queue, ns)
		}
	}

	return -1

}
