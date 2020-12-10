package year2020

import (
	"advent-of-code/read"
)

func getSeat(s string) int {
	topRow := 127
	botRow := 0
	topCol := 7
	botCol := 0

	for _, c := range s {
		switch c {
		case 'F':
			topRow -= (topRow - botRow + 1) / 2
		case 'B':
			botRow += (topRow - botRow + 1) / 2
		case 'L':
			topCol -= (topCol - botCol + 1) / 2
		case 'R':
			botCol += (topCol - botCol + 1) / 2
		}
	}
	return botRow*8 + botCol
}

func solution05a(data []string) int {
	var max int
	for _, s := range data {
		if n := getSeat(s); n > max {
			max = n
		}
	}
	return max
}

func solution05b(data []string) int {
	m := make(map[int]bool)

	for _, s := range data {
		m[getSeat(s)] = true
	}

	for i := 0; i < 999; i++ {
		if m[i-1] && !m[i] && m[i+1] {
			return i
		}
	}
	return -1
}

func solution05(filename string) (int, int) {
	data := read.StringSlice(filename)

	return solution05a(data), solution05b(data)
}
