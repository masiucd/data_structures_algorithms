package linkedlist

import (
	"fmt"
	"strings"
)

type SingleList struct {
	Head *SingleNode
	Tail *SingleNode
	Size int
}

func (l *SingleList) Append(value int) {
	newNode := &SingleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++

}
func (l *SingleList) Prepend(value int) {
	newNode := &SingleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Size++

}

func (l *SingleList) InsertAt(value, index int) {
	if l.Head == nil {
		l.Append(value)
	} else {
		newNode := &SingleNode{Value: value}
		prevNode := l.Get(index - 1)
		nextNode := prevNode.Next
		prevNode.Next = newNode
		newNode.Next = nextNode
		l.Size++
	}

}
func (l *SingleList) Search(value int) bool {
	if l.Head == nil || l.Tail == nil {
		return false
	}
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *SingleList) Get(index int) *SingleNode {
	if index < 0 || index >= l.Size {
		return nil
	}

	current := l.Head
	for counter := 0; counter < index; counter++ {
		current = current.Next
	}

	return current
}

func (l *SingleList) Delete(value int) {

}

func (l *SingleList) Traverse() {}

func Reverse() {}

func (l *SingleList) Print() {
	current := l.Head
	var xs []string
	for current != nil {
		xs = append(xs, fmt.Sprintf("%d", current.Value))
		current = current.Next
	}
	fmt.Println(strings.Join(xs, " -> "))
}

func New() *SingleList {
	return &SingleList{}
}
