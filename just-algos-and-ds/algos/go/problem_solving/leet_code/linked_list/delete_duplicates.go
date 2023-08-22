package linked_list

func deleteDuplicates(head *ListNode) *ListNode {
	n := head
	for n != nil {
		if n.Next == nil {
			break
		}
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return head
}
