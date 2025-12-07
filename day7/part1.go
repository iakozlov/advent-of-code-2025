package day7

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

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	firstLine := lines[0]
	beams := make([]bool, len(firstLine))
	for i := range firstLine {
		if firstLine[i] == 'S' {
			beams[i] = true
			break
		}
	}

	splits := 0

	for _, row := range lines[1:] {
		prev := append([]bool(nil), beams...)

		for i := 0; i < len(row); i++ {
			if prev[i] && row[i] == '^' {
				splits++
				beams[i] = false
				if i > 0 {
					beams[i-1] = true
				}
				if i+1 < len(beams) {
					beams[i+1] = true
				}
			}
		}
	}

	return splits, nil
}
