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
		// a := list.Tail
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
		//
	}
	list.Size++
}

func (list *DoubleList) Prepend(value int) {
	//
}

func (list *DoubleList) InsertAt(value, index int) {
	//
}

func (list *DoubleList) Search(value int) bool {
	return false
}

func (list *DoubleList) Get(index int) *linkedlist.DoubleNode {
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
