package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifference(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	got := findDifference(nums1, nums2)
	want := [][]int{{1, 3}, {4, 6}}
	assert.Equal(t, want, got)

}

func TestFindDifference2(t *testing.T) {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	got := findDifference(nums1, nums2)
	want := [][]int{{3}, {}}
	if got[0][0] != 3 && len(got[1]) != 0 {
		t.Errorf("got %v want %v", got, want)
	}

}
