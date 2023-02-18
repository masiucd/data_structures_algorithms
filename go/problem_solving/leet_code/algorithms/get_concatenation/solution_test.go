package getconcatenation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConcatenation(t *testing.T) {

	t.Run("Test Case 1", func(t *testing.T) {
		got := getConcatenation([]int{1, 2, 1})
		want := []int{1, 2, 1, 1, 2, 1}
		assert.Equal(t, got, want)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		got := getConcatenation([]int{1, 3, 2, 1})
		want := []int{1, 3, 2, 1, 1, 3, 2, 1}
		assert.Equal(t, got, want)
	})

}
