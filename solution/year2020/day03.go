package year2020

import (
	"advent-of-code/read"
)

type slope struct {
	x int
	y int
}

func (s slope) run(data []string) int {
	var count int
	var x, y int

	for y < len(data) {
		if data[y][x] == '#' {
			count++
		}

		x = (x + s.x) % (len(data[y]))
		y += s.y
	}

	return count
}

func solution03a(data []string) int {
	s := slope{3, 1}
	return s.run(data)
}

func solution03b(data []string) int {
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	mul := 1
	for _, s := range slopes {
		mul *= s.run(data)
	}

	return mul
}

func solution03(filename string) (int, int) {
	data := read.StringSlice(filename)

	return solution03a(data), solution03b(data)
}
