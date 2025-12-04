package day4

func canBeAccessed(i int, j int, grid [][]string) bool {
	positions := [][2]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}
	countAdjacentRolls := 0

	for _, pos := range positions {
		if isPositionRoll(pos[0], pos[1], grid) {
			countAdjacentRolls++
		}
	}

	return countAdjacentRolls < 4
}
func isPositionRoll(i int, j int, grid [][]string) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}
	return grid[i][j] == "@"
}
