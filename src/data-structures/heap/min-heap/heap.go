package minheap

import "fmt"

type Heap struct {
	heap []int
}

func NewHeap() *Heap {
	return &Heap{
		heap: []int{0},
	}
}

func (h *Heap) Push(val int) {
	h.heap = append(h.heap, val)
	i := len(h.heap) - 1

	// Percolate up
	for i > 1 && h.heap[i] < h.heap[i/2] {
		// Swap values
		tmp := h.heap[i]
		h.heap[i] = h.heap[i/2]
		h.heap[i/2] = tmp
		i = i / 2
	}
}

func (h *Heap) Print() {
	fmt.Println(h.heap)
}

func (h *Heap) Top() int {
	return h.heap[1]
}

func (h *Heap) Size() int {
	return len(h.heap) - 1
}

func (h *Heap) Pop() int {
	if len(h.heap) == 0 {
		return -1
	}
	if len(h.heap) == 2 {
		last := h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]
		return last
	}
	res := h.heap[1]
	// move last value to the root
	h.heap[1] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	i := 1

	// Propagate down
	// while we have a left child
	for {
		leftChild := 2 * i
		rightChild := 2*i + 1
		// ensure we are within bounds
		if rightChild < len(h.heap) {
			if h.heap[rightChild] < h.heap[leftChild] && h.heap[i] > h.heap[rightChild] {
				// swap the right child with the parent
				h.heap[i], h.heap[rightChild] = h.heap[rightChild], h.heap[i]
				i = rightChild
			} else if h.heap[i] > h.heap[leftChild] {
				// swap the left child with the parent
				h.heap[i], h.heap[leftChild] = h.heap[leftChild], h.heap[i]
				i = leftChild
			} else {
				break
			}
		} else if leftChild < len(h.heap) {
			if h.heap[i] > h.heap[leftChild] {
				// swap the left child with the parent
				h.heap[i], h.heap[leftChild] = h.heap[leftChild], h.heap[i]
				i = leftChild
			} else {
				break
			}
		} else {
			break
		}
	}
	return res
}

// Formula for heapify - only works for complete binary trees
// leftChild = 2 * i - always even since 2 * x is always even
// rightChild = 2 * i + 1 - always odd since 2 * x + 1 is always odd
// parent = i / 2
type Heapifyable interface {
	Top() int
	Pop() int
	Push(int)
	Print()
	Size() int
}
