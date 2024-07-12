package searchbst

import (
	"go-ds/src/data-structures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBst(t *testing.T) {
	root := trees.TreeNode{Val: 4}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 7}
	root.Left.Left = &trees.TreeNode{Val: 1}
	root.Left.Right = &trees.TreeNode{Val: 3}
	got := searchBST(&root, 2)
	want := []int{1, 2, 3}
	assert.Equal(t, want, nodesToList(got))

	got = searchBST(&root, 5)
	want = []int{}
	assert.Equal(t, want, nodesToList(got))

	got = searchBST(&root, 4)
	want = []int{
		1, 2, 3, 4, 7,
	}
	assert.Equal(t, want, nodesToList(got))

}

func nodesToList(root *trees.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(nodesToList(root.Left), root.Val), nodesToList(root.Right)...)
}
