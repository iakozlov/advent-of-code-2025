package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var ranges []IdRange

	freshIdsCount := 0

	isReadingIds := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isReadingIds = true
			ranges = MergeIntervals(ranges)
			continue
		}
		if !isReadingIds {
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
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}

			if InRangeSearch(id, ranges) {
				freshIdsCount++
			}
		}
	}

	return freshIdsCount, nil
}
