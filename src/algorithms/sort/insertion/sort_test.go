package insertion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	got := insertionSort([]int{1, 22, 3, 4, 1})
	want := []int{1, 1, 3, 4, 22}

	assert.Equal(t, got, want)
}
