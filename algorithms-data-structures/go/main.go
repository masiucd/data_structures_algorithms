package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree interface {
	Insert(value int)
	Search(value int) bool
	Remove(value int)
	Size() int
	Traverse()
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
		t.root = &Node{value: value}
		return
	}
	insert(t.root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	}
	if value < node.value {
		node.left = insert(node.left, value)
	} else if value > node.value {
		node.right = insert(node.right, value)
	}
	return node
}

func (t *Bst) Search(value int) bool { return false }
func (t *Bst) Remove(value int)      {}
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
	return 1 + size(node.left) + size(node.right)
}
func (t *Bst) Traverse()               {}
func (t *Bst) Find(value int) *Node    { return nil }
func (t *Bst) Contains(value int) bool { return false }

func (t *Bst) BFS() []int {
	var result []int
	q := []Node{*t.root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		result = append(result, node.value)
		if node.left != nil {
			q = append(q, *node.left)
		}
		if node.right != nil {
			q = append(q, *node.right)
		}
	}
	return result
}

func (t *Bst) DFS(order string) []int { return nil }
func inOrder(result *[]int)           {}
func preOrder(result *[]int)          {}
func postOrder(result *[]int)         {}

func main() {
	bst := NewBst()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(2)
	bst.Insert(1)
	fmt.Println(bst.BFS())
	fmt.Println(bst.Size())

}
