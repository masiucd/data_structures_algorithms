package mergelist

// ListNode is a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newList *ListNode = &ListNode{}
	var tail = newList

	for list1 != nil && list2 != nil {
		// build the list backwords!!
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return newList.Next
}
