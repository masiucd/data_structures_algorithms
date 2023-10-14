package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	res := mergeTwoLists(list1, list2)

	r := printList(res)
	expected := []int{1, 1, 2, 3, 4, 4}

	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], r[i])

	}
}

func printList(list *ListNode) []int {
	var xs []int
	for list != nil {
		xs = append(xs, list.Val)
		list = list.Next
	}
	return xs
}
