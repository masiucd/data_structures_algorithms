package mergelist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeList(t *testing.T) {
	got := mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}})
	want := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}
	assert.Equal(t, got, want, "Expected %v  got %v", want, got)
}
