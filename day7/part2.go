package day7

import (
	"bufio"
	"os"
)

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	firstLine := lines[0]
	var startColumn int
	for i := 0; i < len(firstLine); i++ {
		if firstLine[i] == 'S' {
			startColumn = i
			break
		}
	}

	timelinesCount := make([]int, len(firstLine))
	timelinesCount[startColumn] = 1

	for _, row := range lines[1:] {
		next := make([]int, len(row))

		for i, value := range timelinesCount {
			if value == 0 {
				continue
			}
			if row[i] == '^' {
				if i-1 >= 0 {
					next[i-1] += value
				}
				if i+1 < len(row) {
					next[i+1] += value
				}
			} else {
				next[i] += value
			}
		}

		timelinesCount = next
	}

	result := 0
	for _, value := range timelinesCount {
		result += value
	}

	return result, nil
}
