package binary_search_tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Bst struct {
	Root *Node
}

func NewBst() *Bst {
	return &Bst{
		Root: nil,
	}
}

func (bst *Bst) Insert(key int) {
	newNode := Node{Value: key}
	if bst.Root == nil {
		bst.Root = &newNode
		bst.Root.Value = key
		bst.Root.Left = nil
		bst.Root.Right = nil

	} else {
		current := bst.Root
		for {
			if key < current.Value {
				//	go left
				if current.Left != nil {
					current = current.Left
				} else {
					current.Left = &newNode
					return
				}
			} else if key > current.Value {
				//	Go right
				if current.Right != nil {
					current = current.Right
				} else {
					current.Right = &newNode
					return
				}
			}
		}
	}
}

func (bst *Bst) InsertRec(key int) {
	bst.Root = insertHelper(bst.Root, key)
}
func insertHelper(root *Node, key int) *Node {
	if root == nil {
		root = &Node{Value: key}
	}
	if key < root.Value {
		//	go left
		root.Left = insertHelper(root.Left, key)
	} else if key > root.Value {
		//    go right
		root.Right = insertHelper(root.Right, key)
	}
	return root
}

func (bst *Bst) Search(key int) *Node {
	return searchHelper(bst.Root, key)
}

func searchHelper(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key > root.Value {
		return searchHelper(root.Right, key)
	} else if key < root.Value {
		return searchHelper(root.Left, key)
	} else {
		return root
	}
}

func (bst *Bst) Remove(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key > root.Value {
		root.Right = bst.Remove(root.Right, key)
	} else if key < root.Value {
		root.Left = bst.Remove(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			var minNode = bst.MinValue(root.Right)
			root.Value = minNode.Value
			root.Right = bst.Remove(root.Right, minNode.Value)
		}
	}
	return root
}

func (bst *Bst) MinValue(root *Node) *Node {
	var current = root
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *Bst) MaxValue(root *Node) *Node {
	var current = root
	for current != nil && current.Right != nil {
		current = current.Right
	}
	return current
}

func (bst *Bst) Inorder() []int {
	var result []int
	var fn func(*Node)

	fn = func(root *Node) {
		if root == nil {
			return
		}
		fn(root.Left)
		result = append(result, root.Value)
		fn(root.Right)
	}

	fn(bst.Root)
	return result
}

func (bst *Bst) Preorder() []int {
	var result []int
	var fn func(*Node)

	fn = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Value)
		fn(root.Left)
		fn(root.Right)
	}

	fn(bst.Root)
	return result
}

func (bst *Bst) Postorder() []int {
	var result []int
	var fn func(*Node)

	fn = func(root *Node) {
		if root == nil {
			return
		}
		fn(root.Left)
		fn(root.Right)
		result = append(result, root.Value)
	}

	fn(bst.Root)
	return result
}

func (bst *Bst) BreadthFirstSearch() []int {
	result := []int{}
	queue := []*Node{bst.Root}
	level := 0
	for len(queue) > 0 {
		fmt.Println("Level ", level)
		for range queue {
			current := queue[0]
			result = append(result, current.Value)
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		level++
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (bst *Bst) MaxDepth(root *Node) int {
	if bst.Root == nil {
		return 0
	}
	return max(bst.MaxDepth(bst.Root.Left), bst.MaxDepth(bst.Root.Left)) + 1
}
