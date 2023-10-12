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

}
func (l *SingleList) InsertAt(value, index int) {

}
func (l *SingleList) Search(value int) bool {
	return false
}

func (l *SingleList) Get(index int) int {
	return 0
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
