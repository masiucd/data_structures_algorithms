package sortbyheight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByHeight(t *testing.T) {
	tests := []struct {
		heights []int
		want    []int
	}{
		{[]int{-1, 150, 190, 170, -1, -1, 160, 180}, []int{-1, 150, 160, 170, -1, -1, 180, 190}},
		{[]int{-1, -1, -1, -1, -1}, []int{-1, -1, -1, -1, -1}},
		{[]int{4, 2, 9, 11, 2, 16}, []int{2, 2, 4, 9, 11, 16}},
		{[]int{4, 2, 9, 11, 2, 16, 1, 3}, []int{1, 2, 2, 3, 4, 9, 11, 16}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, sortByHeight(tt.heights))
	}
}
