package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part1(filename string) (int, error) {
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
			currPosition = (currPosition - steps) % 100
			if currPosition < 0 {
				currPosition += 100
			}
		} else {
			// right rotation
			currPosition = (currPosition + steps) % 100
		}
		if currPosition == 0 {
			numOfZeroes++
		}
	}
	return numOfZeroes, nil
}
