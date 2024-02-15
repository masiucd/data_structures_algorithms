package aresimilar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreSimilar(t *testing.T) {
	tests := []struct {
		a, b []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{2, 1, 3}, true},
		{[]int{1, 2, 2}, []int{2, 1, 1}, false},
		{[]int{1, 1, 4}, []int{1, 2, 3}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, areSimilar(test.a, test.b))
	}
}
