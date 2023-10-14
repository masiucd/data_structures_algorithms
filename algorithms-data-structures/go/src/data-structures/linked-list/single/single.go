package single

import (
	"fmt"
	"go-ds/src/data-structures/linked-list"
	"strings"
)

type SingleList struct {
	Head *linkedlist.linkedlist
	Tail *linkedlist.SingleNode
	Size int
}

func (l *SingleList) Append(value int) {
	newNode := &linkedlist.SingleNode{Value: value}
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
	newNode := &linkedlist.SingleNode{Value: value}
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
		newNode := &linkedlist.SingleNode{Value: value}
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

func (l *SingleList) Get(index int) *linkedlist.SingleNode {
	if index < 0 || index >= l.Size {
		return nil
	}

	current := l.Head
	for counter := 0; counter < index; counter++ {
		current = current.Next
	}

	return current
}

func (l *SingleList) Delete(index int) {
	if index < 0 || index >= l.Size {
		return
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Tail = l.Head
		l.Size--
		return
	}

	prevNode := l.Get(index - 1)
	if index == l.Size-1 {
		l.Tail = prevNode
		prevNode.Next = nil
		l.Size--
		return
	}

	prevNode.Next = prevNode.Next.Next
	l.Size--

}

func (l *SingleList) Reverse() {
	if l.Size <= 1 {
		return
	}
	current := l.Head
	var prev *linkedlist.SingleNode
	var next *linkedlist.SingleNode
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Tail = l.Head
	l.Head = prev

}

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
