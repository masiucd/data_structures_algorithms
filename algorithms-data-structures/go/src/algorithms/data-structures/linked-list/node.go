package linkedlist

type SingleNode struct {
	Value int
	Next  *SingleNode
}

type DoubleNode struct {
	Value int
	Next  *DoubleNode
	Prev  *DoubleNode
}

type Listable interface {
	Append(value int)
	Prepend(value int)
	InsertAt(value, index int)
	Search(value int) bool
	Get(index int) *SingleNode
	Delete(value int)
	Traverse()
	Reverse()
	Print()
}
