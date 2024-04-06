package countdigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAmountOfDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{10, 2},
		{100, 3},
		{1000, 4},
		{10000, 5},
		{100000, 6},
		{1000000, 7},
		{10000000, 8},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, countAmountOfDigits(tt.n))
	}
}
