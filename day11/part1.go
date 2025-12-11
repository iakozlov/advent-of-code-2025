package day11

import (
	"log"
)

func Part1(filename string) (int, error) {
	adj, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	visited := make(map[string]bool)
	memo := make(map[string]int)
	result := dfsCountPaths("you", "out", adj, visited, memo)

	return result, nil
}

func dfsCountPaths(node string, targetNode string, adj map[string][]string, visited map[string]bool, memo map[string]int) int {
	if val, ok := memo[node]; ok {
		return val
	}
	if visited[node] {
		log.Fatal("Cycle detected")
	}
	visited[node] = true
	defer delete(visited, node)
	if node == targetNode {
		memo[node] = 1
		return 1
	}
	var total int
	for _, dst := range adj[node] {
		total += dfsCountPaths(dst, targetNode, adj, visited, memo)
	}
	memo[node] = total
	return total
}
