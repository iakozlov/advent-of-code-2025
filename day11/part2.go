package day11

import (
	"log"
)

type state struct {
	node   string
	hasDAC bool
	hasFFT bool
}

func Part2(filename string) (int, error) {
	adj, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	visited := make(map[string]bool)
	memo := make(map[state]int)

	paths := dfsCountPathsDacFft("svr", "out", adj, visited, memo, false, false)

	return paths, nil

}

func dfsCountPathsDacFft(node string, target string, adj map[string][]string, visited map[string]bool, memo map[state]int, hasDAC bool, hasFFT bool) int {
	if visited[node] {
		log.Fatal("Cycle detected")
	}
	visited[node] = true
	defer delete(visited, node)

	if node == "dac" {
		hasDAC = true
	}
	if node == "fft" {
		hasFFT = true
	}

	currentState := state{node: node, hasDAC: hasDAC, hasFFT: hasFFT}

	if val, ok := memo[currentState]; ok {
		return val
	}

	if node == target {
		result := 0
		if hasFFT && hasDAC {
			result = 1
		}
		memo[currentState] = result
		return result
	}

	totalPaths := 0
	for _, dst := range adj[node] {
		totalPaths += dfsCountPathsDacFft(dst, target, adj, visited, memo, hasDAC, hasFFT)
	}

	memo[currentState] = totalPaths
	return totalPaths
}
