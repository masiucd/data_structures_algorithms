package simplequeue

type QueueNode struct {
	Val  int
	next *QueueNode
}

type Queue struct {
	Head *QueueNode
	Size int
}

func NewQ() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val int) {
	newNode := QueueNode{Val: val}
	if q.Head == nil {
		q.Head = &newNode

	} else {
		current := q.Head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
	q.Size++
}

func (q *Queue) Pop() *QueueNode {
	if q.Head == nil {
		return nil
	}
	current := q.Head
	q.Head = current.next
	current.next = nil
	return current
}

func (q *Queue) Peek() int {
	return q.Head.Val
}

func (q *Queue) Print() []int {
	var res []int
	current := q.Head
	for current != nil {
		res = append(res, current.Val)
		current = current.next
	}
	return res
}
