package linked_list

//https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := head
	if head.Next != nil {
		newHead = reverseListRecursive(head.Next)
		//Reverse link between next node and head
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}
