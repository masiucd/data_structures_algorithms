package stack

// MinStack struct
type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor for MinStack
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}

}

// Push pushes the stack
func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	var size int = len(ms.minStack)
	if size == 0 {
		ms.minStack = append(ms.minStack, val)
	} else if size > 0 {
		if val < ms.minStack[size-1] {
			ms.minStack = append(ms.minStack, val)
		} else {
			ms.minStack = append(ms.minStack, ms.minStack[size-1])
		}
	}
}

// Pop pop from the stack
func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

// Top returns the top of the stack
func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

// GetMin returns the min value
func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
