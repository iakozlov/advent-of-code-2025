package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2(filename string) (int, error) {
	var numOfZeroes = 0
	var currPosition = 50
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation := scanner.Text()
		steps, err := strconv.Atoi(rotation[1:])
		if err != nil {
			return 0, err
		}

		if strings.HasPrefix(rotation, "L") {
			// left rotation
			numOfZeroes += steps / 100
			if currPosition > 0 && steps%100 >= currPosition {
				numOfZeroes++
			}
			currPosition = (currPosition - steps) % 100
			if currPosition < 0 {
				currPosition += 100
			}
		} else {
			// right rotation
			numOfZeroes += (currPosition + steps) / 100
			currPosition = (currPosition + steps) % 100
		}
	}
	return numOfZeroes, nil
}
