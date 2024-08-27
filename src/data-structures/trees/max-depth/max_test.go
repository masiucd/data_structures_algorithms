package maxdepth

import (
	"go-ds/src/data-structures/trees"
	"testing"
)

var tests = []struct {
	name string
	tree *trees.TreeNode
	want int
}{
	{
		name: "Test case 1",
		tree: &trees.TreeNode{Val: 3,
			Left: &trees.TreeNode{Val: 9},
			Right: &trees.TreeNode{Val: 20,
				Left:  &trees.TreeNode{Val: 15},
				Right: &trees.TreeNode{Val: 7},
			},
		},
		want: 3,
	},
	{
		name: "Test case 2",
		tree: &trees.TreeNode{Val: 1,
			Left: &trees.TreeNode{Val: 2,
				Left: &trees.TreeNode{Val: 4,
					Left: &trees.TreeNode{Val: 6,
						Left:  &trees.TreeNode{Val: 8},
						Right: &trees.TreeNode{Val: 9},
					},
					Right: &trees.TreeNode{Val: 7,
						Left:  &trees.TreeNode{Val: 10},
						Right: &trees.TreeNode{Val: 11},
					},
				},
				Right: &trees.TreeNode{Val: 5,
					Left:  &trees.TreeNode{Val: 12},
					Right: &trees.TreeNode{Val: 13},
				},
			},
			Right: &trees.TreeNode{Val: 3,
				Left: &trees.TreeNode{Val: 14},
			},
		},
		want: 5,
	},
}

func TestMaxDepth(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.tree); got != tt.want {
				t.Errorf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthV2(tt.tree); got != tt.want {
				t.Errorf("maxDepthV2() = %d, want %d", got, tt.want)
			}
		})
	}
}
