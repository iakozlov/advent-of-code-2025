package main

import (
	"advent-of-code-2025/day7"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: day7 <part>")
		fmt.Println("  part: 1 or 2")
		os.Exit(1)
	}

	part := os.Args[1]
	inputFile := "day7/input/input.txt"

	switch part {
	case "1":
		result, err := day7.Part1(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Part 1 Result: %d\n", result)
	case "2":
		result, err := day7.Part2(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Part 2 Result: %d\n", result)
	default:
		fmt.Fprintf(os.Stderr, "Invalid part: %s\n", part)
		os.Exit(1)
	}
}
