package minstack

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	if len(ms.stack) == 0 {
		ms.stack = append(ms.stack, val)
		ms.mins = append(ms.mins, val)
	} else {
		top := ms.mins[len(ms.mins)-1]
		if val < top {
			ms.mins = append(ms.mins, val)
		} else {
			ms.mins = append(ms.mins, top)
		}
		ms.stack = append(ms.stack, val)
	}

}

func (ms *MinStack) Pop() {
	if len(ms.stack) > 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
		ms.mins = ms.mins[:len(ms.mins)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) > 0 {
		return ms.stack[len(ms.stack)-1]
	}
	return -1
}

func (ms *MinStack) GetMin() int {
	if len(ms.mins) == 0 {
		return -1
	}
	return ms.mins[len(ms.mins)-1]
}
