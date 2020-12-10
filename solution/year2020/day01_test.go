package year2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution01a(t *testing.T) {
	give := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want := 514579

	assert.Equal(t, want, solution01a(give))
}

func TestSolution01b(t *testing.T) {
	give := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	want := 241861950

	assert.Equal(t, want, solution01b(give))
}
