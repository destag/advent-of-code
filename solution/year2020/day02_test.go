package year2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution02a(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			assert.Equal(t, tt.want, solution02a(tt.give))
		})
	}
}

func TestSolution02b(t *testing.T) {
	tests := []struct {
		give string
		want bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			assert.Equal(t, tt.want, solution02b(tt.give))
		})
	}
}
