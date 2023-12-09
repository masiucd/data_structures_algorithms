package bst

import (
	ll "github.com/bahlo/generic-list-go"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree interface {
	Insert(value int)
	Remove(value int)
	Size() int
	Find(value int) *Node
	Contains(value int) bool
	BFS() []int
	DFS(order string) []int // in, pre, post

}

func NewBst() Tree {
	return &Bst{}
}

type Bst struct {
	root *Node
}

func (t *Bst) Insert(value int) {
	if t.root == nil {
		t.root = &Node{Value: value}
		return
	}
	insert(t.root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (t *Bst) Remove(value int) {
	if t.root == nil {
		return
	}

}

func (t *Bst) Size() int {
	if t.root == nil {
		return 0
	}
	return size(t.root)
}
func size(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.Left) + size(node.Right)
}

func (t *Bst) Find(value int) *Node {
	if t.root == nil {
		return nil
	}
	return find(t.root, value)
}

func find(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value == node.Value {
		return node
	}
	find(node.Left, value)
	find(node.Right, value)
	return nil
}

func (t *Bst) Contains(value int) bool {
	return t.Find(value) != nil
}

func (t *Bst) BFS() []int {
	var result []int
	queue := ll.New[*Node]()
	queue.PushBack(t.root)
	for queue.Len() > 0 {
		element := queue.Front()
		node := element.Value
		queue.Remove(element)
		result = append(result, node.Value)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}

func (t *Bst) DFS(order string) []int {
	var result []int
	if t.root == nil {
		return result
	}
	switch order {
	case "IN":
		inOrder(t.root, &result)
	case "PRE":
		preOrder(t.root, &result)
	case "POST":
		postOrder(t.root, &result)
	}
	return result
}
func inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Value)
	inOrder(node.Right, result)
}
func preOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Value)
	preOrder(node.Left, result)
	preOrder(node.Right, result)
}
func postOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, result)
	postOrder(node.Right, result)
	*result = append(*result, node.Value)
}
