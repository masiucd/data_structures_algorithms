package can_reach_leaf

import (
	"data_structures_algos_go/common/tree"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanReachLeaf(t *testing.T) {
	node := tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val:   0,
			Right: &tree.TreeNode{Val: 7},
		},
		Right: &tree.TreeNode{
			Val:   1,
			Left:  &tree.TreeNode{Val: 2},
			Right: &tree.TreeNode{Val: 0},
		},
	}

	got := canReachLeaf(&node)
	assert.Equal(t, got, true, fmt.Sprintf("Expected %v, got %v", got, true))

	node = tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val:   0,
			Right: &tree.TreeNode{Val: 7},
		},
		Right: &tree.TreeNode{
			Val:   1,
			Left:  &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 0}},
			Right: &tree.TreeNode{Val: 0},
		},
	}

	got = canReachLeaf(&node)
	assert.Equal(t, got, false, fmt.Sprintf("Expected %v, got %v", got, false))
}

func TestCanReachLeafWithPath(t *testing.T) {
	node := tree.TreeNode{
		Val: 4,
		Left: &tree.TreeNode{
			Val:   0,
			Right: &tree.TreeNode{Val: 7},
		},
		Right: &tree.TreeNode{
			Val:   1,
			Left:  &tree.TreeNode{Val: 3, Right: &tree.TreeNode{Val: 0}},
			Right: &tree.TreeNode{Val: 2},
		},
	}

	var paths []int
	expectedPaths := []int{4, 1, 2}
	got := canReachLeafWithPath(&node, &paths)
	assert.Equal(t, got, true, fmt.Sprintf("Expected %v, got %v", got, true))
	for i := range paths {
		if paths[i] != expectedPaths[i] {
			t.Errorf("Expected %v, got %v", expectedPaths[i], paths[i])
		}
	}

}
