package year2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution03a(t *testing.T) {
	give := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	want := 7

	assert.Equal(t, want, solution03a(give))
}

func TestSolution03b(t *testing.T) {
	give := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	want := 336

	assert.Equal(t, want, solution03b(give))
}
