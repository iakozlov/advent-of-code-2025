package day3

import "strconv"

func getMaxNumberInString(line string, idxShift int) (int, int, error) {
	var maxNumber = 0
	var maxNumberIdx = 0
	for i := 0; i < len(line); i++ {
		char := line[i]
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, 0, err
		}

		if digit > maxNumber {
			maxNumber = digit
			maxNumberIdx = i
		}
	}
	return maxNumber, maxNumberIdx + idxShift, nil
}
