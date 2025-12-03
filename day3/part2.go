package day3

import (
	"bufio"
	"os"
	"strconv"
)

func Part2(filename string) (int, error) {
	var joltageSum = 0
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		result, err := getJoltageOfTwelveDigitBank(bank)
		if err != nil {
			return 0, err
		}
		joltageSum += result
	}

	return joltageSum, nil
}

func getJoltageOfTwelveDigitBank(bank string) (int, error) {
	resultString := ""
	var prevIndexToStartWith = 0
	for i := 0; i < 12; i++ {
		digit, idx, err := getMaxNumberInString(bank[prevIndexToStartWith:len(bank)-(12-i-1)], prevIndexToStartWith)
		if err != nil {
			return 0, err
		}
		resultString += strconv.Itoa(digit)
		prevIndexToStartWith = idx + 1
	}

	result, err := strconv.Atoi(resultString)
	if err != nil {
		return 0, err
	}

	return result, nil
}
