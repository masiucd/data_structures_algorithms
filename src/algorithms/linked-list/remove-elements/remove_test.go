package removeelements

import (
	"fmt"
	linkedlist "go-ds/src/data-structures/linked-list"
	"testing"
)

func buildLinkedList(values []int) *linkedlist.SingleNode {
	var head *linkedlist.SingleNode
	var current *linkedlist.SingleNode

	for _, value := range values {
		node := &linkedlist.SingleNode{Value: value}
		if head == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = node
		}
	}

	return head
}

func linkedListToSlice(head *linkedlist.SingleNode) []int {
	var values []int
	current := head

	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}

	return values
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		input    []int
		remove   int
		expected []int
	}{
		{
			input:    []int{7, 7, 7, 7, 7, 7, 7},
			remove:   7,
			expected: []int{},
		},
		{
			input:    []int{1, 2, 6, 3, 4, 5, 6},
			remove:   6,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		head := buildLinkedList(test.input)
		res := removeElements(head, test.remove)
		got := linkedListToSlice(res)

		fmt.Println(got, test.expected)
		if len(got) != len(test.expected) || !slicesEqual(got, test.expected) {
			t.Errorf("Got %v, expected %v", got, test.expected)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
