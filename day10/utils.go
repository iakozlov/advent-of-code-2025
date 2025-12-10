package day10

import (
	"strconv"
	"strings"
)

func parseLinePart1(line string) (uint64, []uint64, error) {
	var (
		targetMask uint64
		buttons    []uint64
	)

	start := strings.IndexByte(line, '[')
	end := strings.IndexByte(line, ']')
	lights := line[start+1 : end]
	nLights := len(lights)
	for i, ch := range lights {
		if ch == '#' {
			targetMask |= 1 << uint(i)
		}
	}

	for i := end + 1; i < len(line); i++ {
		if line[i] == '(' {
			j := strings.IndexByte(line[i+1:], ')')
			j += i + 1

			content := strings.TrimSpace(line[i+1 : j])
			if content != "" {
				parts := strings.Split(content, ",")
				var mask uint64
				for _, p := range parts {
					idx, err := strconv.Atoi(p)
					if err != nil {
						return 0, nil, err
					}
					if idx < 0 || idx >= nLights {
						return 0, nil, err
					}
					mask |= 1 << uint(idx)
				}
				buttons = append(buttons, mask)
			}
			i = j
		}
	}

	return targetMask, buttons, nil
}

func parseLinePart2(line string) ([]int, [][]int, error) {
	left := strings.IndexByte(line, '{')
	right := strings.IndexByte(line, '}')
	joltageStr := line[left+1 : right]
	parts := strings.Split(joltageStr, ",")
	var targets []int
	for _, p := range parts {
		v, err := strconv.Atoi(p)
		if err != nil || v < 0 {
			return nil, nil, err
		}
		targets = append(targets, v)
	}

	var buttons [][]int
	for i := 0; i < len(line); i++ {
		if line[i] == '(' {
			j := strings.IndexByte(line[i+1:], ')')
			j += i + 1

			content := line[i+1 : j]
			if content != "" {
				idxParts := strings.Split(content, ",")
				btn := make([]int, 0, len(idxParts))
				seen := make(map[int]struct{}, len(idxParts))
				for _, ip := range idxParts {
					ip = strings.TrimSpace(ip)
					idx, err := strconv.Atoi(ip)
					if err != nil || idx < 0 {
						return nil, nil, err
					}
					if _, ok := seen[idx]; ok {
						continue
					}

					seen[idx] = struct{}{}
					btn = append(btn, idx)
				}
				buttons = append(buttons, btn)
			}
			i = j
		}
	}

	return targets, buttons, nil
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
