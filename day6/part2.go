package day6

import (
	"bufio"
	"os"
	"strings"
)

func Part2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	isDigit := func(ch byte) bool { return ch >= '0' && ch <= '9' }

	var (
		colVals   []int
		colSeen   []bool
		colSepAll []bool
		colOps    []byte
		maxW      int
	)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		fields := strings.Fields(line)
		isOpLine := len(fields) > 0
		for _, tok := range fields {
			if tok != "+" && tok != "*" {
				isOpLine = false
				break
			}
		}

		if len(line) > maxW {
			delta := len(line) - maxW
			colVals = append(colVals, make([]int, delta)...)
			colSeen = append(colSeen, make([]bool, delta)...)
			addSep := make([]bool, delta)
			for i := range addSep {
				addSep[i] = true
			}
			colSepAll = append(colSepAll, addSep...)
			colOps = append(colOps, make([]byte, delta)...)
			maxW = len(line)
		}

		for c := 0; c < len(line); c++ {
			ch := line[c]
			if ch != ' ' {
				colSepAll[c] = false
			}
			if isOpLine {
				if ch == '+' || ch == '*' {
					colOps[c] = ch
				}
				continue
			}
			if isDigit(ch) {
				colSeen[c] = true
				colVals[c] = colVals[c]*10 + int(ch-'0')
			}
		}
	}
	if err := sc.Err(); err != nil {
		return 0, err
	}

	if maxW == 0 {
		return 0, nil
	}

	total := 0

	for col := maxW - 1; col >= 0; {
		if colSepAll[col] {
			col--
			continue
		}
		right := col
		left := col
		for left-1 >= 0 && !colSepAll[left-1] {
			left--
		}
		op := byte(0)
		for j := right; j >= left; j-- {
			if colOps[j] == '+' || colOps[j] == '*' {
				op = colOps[j]
				break
			}
		}
		nums := make([]int, 0, right-left+1)
		for j := left; j <= right; j++ {
			if colSeen[j] {
				nums = append(nums, colVals[j])
			}
		}

		if len(nums) > 0 && (op == '+' || op == '*') {
			if op == '+' {
				sum := 0
				for _, v := range nums {
					sum += v
				}
				total += sum
			} else {
				prod := 1
				for _, v := range nums {
					prod *= v
				}
				total += prod
			}
		}

		col = left - 1
	}

	return total, nil
}
