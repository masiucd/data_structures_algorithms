package queue

import (
	"container/list"
	"fmt"
	"strings"
)

type Queue struct {
	items list.List
}

type Quable interface {
	Enqueue(item int)
	Dequeue() *list.Element
	Peek() *list.Element
	Len() int
	Print() string
}

func (q *Queue) Len() int {
	return q.items.Len()
}

func (q *Queue) Print() string {
	var xs []string
	for e := q.items.Front(); e != nil; e = e.Next() {
		if value, ok := e.Value.(int); ok {
			n := fmt.Sprintf("%d", value)
			xs = append(xs, n)
		}
	}
	return strings.Join(xs, ",")
}

func (q *Queue) Enqueue(item int) {
	q.items.PushBack(item)
}

func (q *Queue) Dequeue() *list.Element {
	node := q.items.Front()
	q.items.Remove(node)
	return node
}

func (q *Queue) Peek() *list.Element {
	return q.items.Front()
}

func New() *Queue {
	return &Queue{}
}
