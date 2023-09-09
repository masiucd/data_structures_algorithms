package linked_list

//https://leetcode.com/problems/design-linked-list/submissions/

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

func newMyLinkedList() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	cur := this.Head
	for i := 0; i < index; i++ {
		if cur == nil {
			return -1
		}
		cur = cur.Next
	}

	if cur == nil {
		return -1
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head == nil {
		this.Head = &Node{Val: val}
	} else {
		this.Head = &Node{
			Val:  val,
			Next: this.Head,
		}
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.Head == nil {
		this.Head = &Node{Val: val}
		return
	}

	cur := this.Head

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = &Node{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	cur := this.Head
	temp := 0

	for cur != nil {
		if temp == index-1 {
			node := Node{
				Val:  val,
				Next: cur.Next,
			}

			cur.Next = &node
			break
		}

		cur = cur.Next
		temp++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || this.Head == nil {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
	}

	cur := this.Head
	temp := 0

	for cur != nil {
		if temp == index-1 && cur.Next != nil {
			cur.Next = cur.Next.Next
			break
		}

		cur = cur.Next
		temp++
	}
}
