package topkfrequent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {

	var tests = []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
	}

	for _, test := range tests {
		got := topKFrequent(test.input, test.k)
		assert.Equal(t, test.expected, got)
	}
}
