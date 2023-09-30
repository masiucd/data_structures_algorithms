package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	printStacks(minStack)
	min := minStack.GetMin() // return -3
	fmt.Println("MIN", min)
	minStack.Pop()
	top := minStack.Top() // return 0
	fmt.Println("top", top)
	min = minStack.GetMin() // return -2
	fmt.Println("MIN", min)

}

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
	return ms.mins[len(ms.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func printStacks(m MinStack) {
	fmt.Println("STACK ", m.stack)
	fmt.Println("MINS ", m.mins)
}
