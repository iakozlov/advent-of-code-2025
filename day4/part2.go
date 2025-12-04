package day4

import (
	"bufio"
	"os"
	"strings"
)

func Part2(filename string) (int, error) {
	grid := [][]string{}
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		row := strings.Split(bank, "")
		grid = append(grid, row)
	}

	result := 0

	for {
		var positionsToUpdate [][2]int
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] == "@" && canBeAccessed(i, j, grid) {
					result++
					positionsToUpdate = append(positionsToUpdate, [2]int{i, j})
				}
			}
		}

		if len(positionsToUpdate) == 0 {
			break
		}

		for _, pos := range positionsToUpdate {
			grid[pos[0]][pos[1]] = "x"
		}
	}

	return result, nil

}
