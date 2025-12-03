package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1(filename string) (int, error) {
	var joltageSum = 0
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		result, err := getJoltageOfBank(bank)
		if err != nil {
			return 0, err
		}
		joltageSum += result
	}

	return joltageSum, nil
}

func getJoltageOfBank(bank string) (int, error) {
	firstDigit, firstDigitIndex, err := getMaxNumberInString(bank[:len(bank)-1], 0)
	if err != nil {
		return 0, err
	}
	secondDigit, _, err := getMaxNumberInString(bank[firstDigitIndex+1:], firstDigitIndex)

	result, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, secondDigit))
	if err != nil {
		return 0, err
	}

	return result, nil
}
