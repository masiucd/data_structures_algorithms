package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorNums(t *testing.T) {
	tt := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "when there is a duplicate and a missing number",
			nums: []int{1, 2, 2, 4},
			want: []int{2, 3},
		},
		{
			name: "when there is a duplicate and a missing number",
			nums: []int{1, 1},
			want: []int{1, 2},
		},
		{
			name: "when there is a duplicate and a missing number",
			nums: []int{2, 2},
			want: []int{2, 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := findErrorNums(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}
}
