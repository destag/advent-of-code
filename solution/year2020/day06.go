package year2020

import (
	"advent-of-code/read"
	"advent-of-code/util"
)

func solution06a(data string) int {
	q := make(map[rune]bool)
	for _, c := range data {
		if c == ':' {
			continue
		}
		q[c] = true
	}

	return len(q)
}

func solution06b(data string) int {
	var group int
	q := make(map[rune]int)
	for _, c := range data {
		if c == ':' {
			group++
			continue
		}
		q[c]++
	}

	var count int
	for _, v := range q {
		if v == group {
			count++
		}
	}

	return count
}

func solution06(filename string) (int, int) {
	data := read.Chunks(filename, ":")

	return util.Sum(data, solution06a), util.Sum(data, solution06b)
}
