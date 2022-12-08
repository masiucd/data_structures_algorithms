package tree

import (
	"data_structures_algos_go/common/tree"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrderOne(t *testing.T) {
	tNode := tree.TreeNode{
		Val:   3,
		Left:  &tree.TreeNode{Val: 9},
		Right: &tree.TreeNode{Val: 20, Left: &tree.TreeNode{Val: 15}, Right: &tree.TreeNode{Val: 7}},
	}
	got := levelOrder(&tNode)
	want := [][]int{{3}, {9, 20}, {15, 7}}
	for i := range got {
		assert.Equal(t, got[i], want[i], fmt.Sprintf("Got %d Want %d", got[i], want[i]))
	}
}
