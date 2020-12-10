package year2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution07a(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{":abc", 3},
		{":a:b:c", 3},
		{"ab:ac", 3},
		{":a:a:a:a", 1},
		{":b", 1},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			assert.Equal(t, tt.want, solution07a(tt.give))
		})
	}
}

func TestSolution07b(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{":abc", 3},
		{":a:b:c", 0},
		{":ab:ac", 1},
		{":a:a:a:a", 1},
		{":b", 1},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			assert.Equal(t, tt.want, solution07b(tt.give))
		})
	}
}
