package countnodes

import (
	"go-ds/src/data-structures/trees"
	"testing"
)

func TestCountNodes(t *testing.T) {

	// Test case 1
	//			1
	//		2		3
	//	4	5	6	7
	// The tree has 7 nodes
	root := &trees.TreeNode{Val: 1}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 3}
	root.Left.Left = &trees.TreeNode{Val: 4}
	root.Left.Right = &trees.TreeNode{Val: 5}
	root.Right.Left = &trees.TreeNode{Val: 6}
	root.Right.Right = &trees.TreeNode{Val: 7}

	expected := 7
	result := CountNodes(root)
	if result != expected {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected, result)
	}
}

// Test case 2
// The tree is empty
func TestCountNodesEmptyTree(t *testing.T) {
	var root *trees.TreeNode
	expected := 0
	result := CountNodes(root)
	if result != expected {
		t.Errorf("Test case 2 failed: expected %v but got%v", expected, result)
	}
}

func BenchmarkCountNodes(b *testing.B) {
	root := &trees.TreeNode{Val: 1}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 3}
	root.Left.Left = &trees.TreeNode{Val: 4}
	root.Left.Right = &trees.TreeNode{Val: 5}
	root.Right.Left = &trees.TreeNode{Val: 6}
	root.Right.Right = &trees.TreeNode{Val: 7}

	for i := 0; i < b.N; i++ {
		CountNodes(root)
	}
}
