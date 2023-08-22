package containsnearbyduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicates(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 1", args{[]int{1, 2, 3, 1}, 3}, true},
		{"Test 2", args{[]int{1, 0, 1, 1}, 1}, true},
		{"Test 3", args{[]int{1, 2, 3, 1, 2, 3}, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solution(tt.args.nums, tt.args.k))
		})
	}
}
