package double

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
	prev *node[T]
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

func (l *List[T]) PushFront(key T) {
	newNode := &node[T]{key: key}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		current := l.head
		newNode.next = current
		current.prev = newNode
		l.head = newNode
	}
	l.size++
}

func (l *List[T]) PushBack(key T) {
	newNode := &node[T]{key: key}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		current := l.tail
		current.next = newNode
		newNode.prev = current
		l.tail = newNode
	}
	l.size++
}

func (l *List[T]) GetAt(index int) *node[T] {
	if index < 0 || index > l.Len() {
		return nil
	}

	switch index {
	case 0:
		return l.head
	case l.Len() - 1:
		return l.tail
	}

	//	check if we should start from head or tail
	var current *node[T]
	var count = 0
	if index < l.Len()/2 {
		//	start from the head
		current = l.head
		for current != nil {
			if count == index {
				return current
			}
			current = current.next
			count++
		}
	} else {
		//	start from the tail
		current = l.tail
		count = l.Len() - 1
		for current != nil {
			if count == index {
				return current
			}
			current = current.prev
			count--
		}
	}
	return l.tail
}

func (l *List[T]) Pop() {
	if l.head == nil {
		return
	} else if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		current := l.tail
		prev := current.prev
		current.prev = nil
		prev.next = nil
		l.tail = prev
	}
	l.size -= 1
}

func (l *List[T]) RemoveStart() {
	if l.head == nil {
		return
	} else if l.Len() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		next := l.head.next
		next.prev = nil
		l.head = next
	}
	l.size -= 1
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) Print() {
	var xs []string
	current := l.head
	for current != nil {
		xs = append(xs, fmt.Sprintf("%v", current.key))
		current = current.next
	}
	fmt.Println(strings.Join(xs, " <---> "))
}

// KeyValue is Nullable
func (n *node[T]) KeyValue() T {
	return n.key
}
