package day6

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

	var numbers [][]int
	var operations []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		if fields[0] == "*" || fields[0] == "+" {
			operations = append(operations, fields...)
			continue
		}

		lineNumbers := make([]int, 0, len(fields))
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return 0, err
			}
			lineNumbers = append(lineNumbers, num)
		}
		numbers = append(numbers, lineNumbers)
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	rows := len(numbers)

	result := 0

	for j, operation := range operations {
		if operation == "+" {
			columnResult := 0
			for i := 0; i < rows; i++ {
				columnResult += numbers[i][j]
			}
			result += columnResult
		} else {
			columnResult := 1
			for i := 0; i < rows; i++ {
				columnResult *= numbers[i][j]
			}
			result += columnResult
		}
	}

	return result, nil
}
