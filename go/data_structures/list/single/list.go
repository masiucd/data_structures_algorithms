package single

import (
	"data_structures_algos_go/data_structures/list"
	"fmt"
	"strings"
)

// Listable actions for the Linked list
type Listable[T list.CustomConstraint] interface {
	PushFront(key T)
	PushBack(key T)
	GetAt(index int) *node[T]
	Pop()
	RemoveStart()
	Len() int
	Reverse()
	Print()
}

type node[T list.CustomConstraint] struct {
	key  T
	next *node[T]
}

type List[T list.CustomConstraint] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewList[T list.CustomConstraint]() *List[T] {
	return &List[T]{head: nil, tail: nil, size: 0}
}

func (l *List[T]) PushBack(key T) {
	newNode := &node[T]{key: key}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *List[T]) PushFront(key T) {
	newNode := &node[T]{key: key}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

func (l *List[T]) GetAt(index int) *node[T] {
	if l.head == nil {
		return nil
	}
	if index < 0 || index >= l.Len() {
		return nil
	}
	if index == 0 {
		return l.head
	}
	if index == l.Len()-1 {
		return l.tail
	}
	counter := 0
	current := l.head
	for current != nil {
		if counter == index {
			return current
		}
		current = current.next
		counter++
	}

	return nil
}

func (l *List[T]) Pop() {
	if l.head == nil {
		return
	}
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		last := l.GetAt(l.Len() - 1)
		lastPrev := l.GetAt(l.Len() - 2)
		if last != nil {
			lastPrev.next = nil
			l.tail = lastPrev
		}
	}
	l.size--

}

func (l *List[T]) RemoveStart() {
	if l.head == nil {
		return
	}
	if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		newHead := l.head.next
		l.head = newHead
	}
	l.size--
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) Reverse() {
	var prev *node[T]
	var current = l.head
	l.head = l.tail
	l.tail = l.head
	for current != nil {
		temp := current.next
		current.next = prev
		prev = current
		current = temp
	}
}

func (l *List[T]) Print() {
	var xs []string
	current := l.head
	for current != nil {
		xs = append(xs, fmt.Sprintf("%v", current.key))
		current = current.next
	}
	fmt.Println(strings.Join(xs, " -> "))
}

// KeyValue is Nullable
func (n *node[T]) KeyValue() T {
	return n.key
}
