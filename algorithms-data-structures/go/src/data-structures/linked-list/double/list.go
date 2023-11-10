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
	//
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
