package double

import (
	linkedlist "go-ds/src/data-structures/linked-list"
)

type DoubleList struct {
	Head *linkedlist.DoubleNode
	Tail *linkedlist.DoubleNode
	Size int
}

func (list *DoubleList) Append(value int) {
	// if list is empty
	node := &linkedlist.DoubleNode{Value: value}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
		//
	}
	list.Size++
}

func (list *DoubleList) Prepend(value int) {
	// if list is empty
	node := &linkedlist.DoubleNode{Value: value}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Prev = node
		list.Head = node
	}
	list.Size++
}

func (list *DoubleList) InsertAt(value, index int) {
	// if list.Head == nil {
	// 	list.Append(value)
	// } else {
	// 	// prevNode :
	// }
}

func (list *DoubleList) Search(value int) bool {
	return false
}

func (list *DoubleList) Get(index int) *linkedlist.DoubleNode {
	if list.Size == 0 || index > list.Size || index < 0 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	if index == list.Size-1 {
		return list.Tail
	}
	middle := (list.Size) / 2
	if index <= middle {
		// start from head
		current := list.Head
		count := 0
		for current != nil {
			if index == count {
				return current
			}
			current = current.Next
			count++
		}
	} else {
		// start from tail
		current := list.Tail
		count := list.Size
		for count >= middle {
			if count == index {
				return current
			}
			current = current.Prev
			count--
		}
	}
	return nil
}

func (list *DoubleList) Delete(index int) {
	//
}

func (list *DoubleList) Reverse() {
}

func (list *DoubleList) Print() {}

func NewDoubleList() *DoubleList {
	return &DoubleList{}
}
