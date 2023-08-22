package dynamic

import "fmt"

type Array struct {
	data     []int
	size     int
	capacity int
}

func newArray(capacity int) *Array {

	data := make([]int, capacity)
	return &Array{data: data, size: 0, capacity: capacity}
}

// O(1)
func (a *Array) push(n int) {
	if a.size == a.capacity {
		a.resize()
	}
	a.data[a.size] = n
	a.size++
}

// O(1)
func (a *Array) pop() int {
	last := a.data[a.size-1]
	a.data = a.data[a.size-1:]
	a.size--
	return last
}

// O(n)
func (a *Array) resize() {
	newCapacity := a.capacity * 2
	a.capacity = newCapacity
	data := make([]int, newCapacity)
	a.data = data

}

func (a *Array) get(index int) int {
	if index < 0 || index > a.size {
		return -1
	}
	item := a.data[index]
	return item
}

func (a *Array) insert(index, n int) {
	if index < a.size || index > a.size {
		return
	}
	a.data[index] = n
}

func (a *Array) print() string {
	var result string
	for _, v := range a.data {
		result += fmt.Sprintf(" %d ", v)
	}
	return result
}
