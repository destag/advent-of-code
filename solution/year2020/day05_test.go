package year2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeat(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			assert.Equal(t, tt.want, getSeat(tt.give))
		})
	}
}
