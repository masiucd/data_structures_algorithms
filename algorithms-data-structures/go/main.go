package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node as the starting point
	dummy := &ListNode{}
	tail := dummy

	// Traverse both lists, merging them
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// If one of the lists is not fully traversed, attach it to the merged list
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

func main() {

	ListOne := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	ListTwo := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	res := mergeTwoLists(ListOne, ListTwo)

	r := printList(res)
	fmt.Println(r)

}

func printList(list *ListNode) []int {
	var xs []int
	for list != nil {
		xs = append(xs, list.Val)
		list = list.Next
	}
	return xs
}
