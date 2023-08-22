package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testHeap() []int {
	heap := NewHeap()
	heap.Push(100)
	heap.Push(50)
	heap.Push(5)
	heap.Push(10)
	return heap.heap
}

func TestMinHeapInsertedCorrectly(t *testing.T) {
	heap := testHeap()
	expected := []int{0, 5, 10, 50, 100}
	assert.Equal(t, expected, heap, fmt.Sprintf("Expected %v actual %v", expected, heap))

}

func TestRootIsTheLowest(t *testing.T) {
	heap := testHeap()
	expected := heap[1]
	rest := heap[2:]
	for _, v := range rest {
		assert.Lessf(t, expected, v, fmt.Sprintf("Expected %v to bee less then %v", expected, v))
	}
}
