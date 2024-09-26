package mixed

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{root, 2, root.Left},
		{root, 5, nil},
	}

	for _, test := range tests {
		got := searchBST(test.root, test.val)
		if got != test.want {
			t.Errorf("searchBST(%v, %d) = %v, want %v", test.root, test.val, got, test.want)
		}
	}
}

func TestSearchBST2(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{root, 2, root.Left},
		{root, 5, nil},
	}

	for _, test := range tests {
		got := searchBST(test.root, test.val)
		if got != test.want {
			t.Errorf("searchBST(%v, %d) = %v, want %v", test.root, test.val, got, test.want)
		}
	}
}
