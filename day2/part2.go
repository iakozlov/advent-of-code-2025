package day2

import (
	"bufio"
	"io"
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
	scanner := bufio.NewReader(file)
	input, err := scanner.ReadString('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}
	input = strings.TrimRight(input, "\r\n")
	ranges := strings.Split(input, ",")
	var result = 0
	for _, idsRange := range ranges {
		idsRange = strings.TrimSpace(idsRange)
		if idsRange == "" {
			continue
		}
		boundaries := strings.Split(idsRange, "-")
		if len(boundaries) != 2 {
			continue
		}
		leftBoundary, err := strconv.Atoi(boundaries[0])
		if err != nil {
			return 0, err
		}
		rightBoundary, err := strconv.Atoi(boundaries[1])
		if err != nil {
			return 0, err
		}
		for id := leftBoundary; id <= rightBoundary; id++ {
			if isWrongIdAtLeastTwice(id) {
				result += id
			}
		}
	}

	return result, nil
}

func isWrongIdAtLeastTwice(id int) bool {
	s := strconv.Itoa(id)
	if len(s) < 2 {
		return false
	}

	for substrLength := 1; substrLength <= len(s)/2; substrLength++ {
		pattern := s[:substrLength]
		numTimesToRepeat := len(s) / substrLength
		if strings.Repeat(pattern, numTimesToRepeat) == s {
			return true
		}
	}
	return false
}
