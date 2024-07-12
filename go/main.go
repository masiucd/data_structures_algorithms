package main

import (
	"fmt"
	linkedlist "go-ds/src/data-structures/linked-list"
)

func main() {
	head := linkedlist.SingleNode{Value: 7}
	head.Next = &linkedlist.SingleNode{Value: 7}
	head.Next.Next = &linkedlist.SingleNode{Value: 7}
	head.Next.Next.Next = &linkedlist.SingleNode{Value: 7}
	head.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 7}
	head.Next.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 7}
	head.Next.Next.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 7}

	// head := linkedlist.SingleNode{Value: 1}
	// head.Next = &linkedlist.SingleNode{Value: 2}
	// head.Next.Next = &linkedlist.SingleNode{Value: 6}
	// head.Next.Next.Next = &linkedlist.SingleNode{Value: 3}
	// head.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 4}
	// head.Next.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 5}
	// head.Next.Next.Next.Next.Next.Next = &linkedlist.SingleNode{Value: 6}

	res := removeElements(&head, 7)
	printList(res)

}

func printList(head *linkedlist.SingleNode) {
	var xs []int
	current := head

	for current != nil {
		xs = append(xs, current.Value)
		current = current.Next
	}

	fmt.Println(xs)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *linkedlist.SingleNode, val int) *linkedlist.SingleNode {
	// Handle the case where the head itself needs to be removed
	for head != nil && head.Value == val {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		if current.Next.Value == val {
			// remove that node
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
