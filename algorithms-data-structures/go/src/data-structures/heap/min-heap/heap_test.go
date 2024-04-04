package minheap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	var h Heapifyable = NewHeap()

	// Push several values
	values := []int{5, 3, 8, 1, 6, 9, 2}
	for _, v := range values {
		h.Push(v)
	}

	// Check that the top value is the smallest one
	if h.Top() != 1 {
		t.Errorf("Top value should be 1, got %d", h.Top())
	}

	// Pop the top value and check that the next top value is the second smallest one
	h.Pop()
	if h.Top() != 2 {
		t.Errorf("Top value should be 2, got %d", h.Top())
	}

	// Continue popping until the heap is empty, each time checking that the top value is the smallest one remaining
	expectedValues := []int{3, 5, 6, 8, 9}
	for _, expected := range expectedValues {
		h.Pop()
		if h.Top() != expected {
			t.Errorf("Top value should be %d, got %d", expected, h.Top())
		}
	}

	// Finally, the heap should be empty
	h.Pop()
	if h.Size() != 0 {
		t.Errorf("Heap should be empty, got %d elements", h.Size())
	}
}
