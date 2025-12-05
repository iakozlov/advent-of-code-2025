package day5

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

	var ranges []IdRange

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		nums := strings.Split(line, "-")
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, err
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, err
		}
		ranges = append(ranges, IdRange{Left: left, Right: right})
	}

	ranges = MergeIntervals(ranges)
	result := 0
	for _, interval := range ranges {
		result += interval.Right - interval.Left + 1
	}

	return result, nil
}
