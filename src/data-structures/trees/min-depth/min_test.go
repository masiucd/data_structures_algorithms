package mindepth

import (
	"go-ds/src/data-structures/trees"
	"testing"
)

var tests = []struct {
	name string
	root *trees.TreeNode
	want int
}{
	{
		name: "Single node",
		root: &trees.TreeNode{Val: 3},
		want: 1,
	},
	{
		name: "Two levels",
		root: &trees.TreeNode{
			Val:   3,
			Left:  &trees.TreeNode{Val: 9},
			Right: &trees.TreeNode{Val: 20},
		},
		want: 2,
	},
	{
		name: "Three levels",
		root: &trees.TreeNode{
			Val:  3,
			Left: &trees.TreeNode{Val: 9},
			Right: &trees.TreeNode{
				Val:   20,
				Left:  &trees.TreeNode{Val: 15},
				Right: &trees.TreeNode{Val: 7},
			},
		},
		want: 2,
	},
	{
		name: "Complex tree",
		root: &trees.TreeNode{
			Val: 2,
			Left: &trees.TreeNode{
				Val:  1,
				Left: &trees.TreeNode{Val: 0},
				Right: &trees.TreeNode{
					Val:   7,
					Left:  &trees.TreeNode{Val: 6},
					Right: &trees.TreeNode{Val: 8},
				},
			},
			Right: &trees.TreeNode{
				Val:  3,
				Left: &trees.TreeNode{Val: 9},
				Right: &trees.TreeNode{
					Val:   1,
					Left:  &trees.TreeNode{Val: 5},
					Right: &trees.TreeNode{Val: 4},
				},
			},
		},
		want: 3,
	},
	{
		name: "Empty tree",
		root: nil,
		want: 0,
	},
}

func TestMinDepth(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDepth(tt.root)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
