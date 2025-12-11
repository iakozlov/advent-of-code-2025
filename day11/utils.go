package day11

import (
	"bufio"
	"os"
	"strings"
)

func parseInput(filename string) (map[string][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	adj := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		from := parts[0]
		dstArr := strings.Split(parts[1], " ")

		for _, dst := range dstArr {
			adj[from] = append(adj[from], dst)
		}
	}

	return adj, nil
}
