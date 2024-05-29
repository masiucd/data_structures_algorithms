package removeelements

import linkedlist "go-ds/src/data-structures/linked-list"

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
