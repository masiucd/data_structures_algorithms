package sumofthetree

import (
	"go-ds/src/data-structures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSum(t *testing.T) {

	var root *trees.TreeNode = &trees.TreeNode{Val: 1,
		Left:  &trees.TreeNode{Val: 2},
		Right: &trees.TreeNode{Val: 3},
	}

	got := treeSum(root)
	want := 6
	assert.Equal(t, got, want)

	root = &trees.TreeNode{Val: 100, Left: &trees.TreeNode{Val: 200}, Right: &trees.TreeNode{Val: 300,
		Left:  &trees.TreeNode{Val: 400, Left: &trees.TreeNode{Val: 500}},
		Right: &trees.TreeNode{Val: 600, Left: &trees.TreeNode{Val: 700, Right: &trees.TreeNode{Val: 800}}},
	}}
	got = treeSum(root)
	want = 3600
	assert.Equal(t, got, want)

}
