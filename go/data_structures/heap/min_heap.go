package heap

// Left child -> i * 2
// Right child -> (i * 2) +1
// Parent -> i/2

type Heap struct {
	heap []int
}

func NewHeap() *Heap {
	return &Heap{
		heap: []int{0},
	}
}

// Push O(log n) runtime
func (h *Heap) Push(val int) {
	h.heap = append(h.heap, val)
	i := len(h.heap) - 1
	parentPosition := i / 2
	parent := h.heap[parentPosition]
	// Percolate up
	for i > 1 && h.heap[i] < parent {
		temp := h.heap[i] // current
		// move the parent into child position
		h.heap[i] = h.heap[parentPosition]
		//move child into parent position
		h.heap[parentPosition] = temp
		//update parent position
		i = i / 2
	}

}

func (h *Heap) Pop() {
	//	TODO implement!

	// Percolate down
}
