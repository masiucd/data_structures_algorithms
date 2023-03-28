function removeDuplicatesFromSortedList(head: ListNode) {
  let current: ListNode | null = head
  while (current) {
    while (current.next !== null && current.next.val === current.val) {
      current.next = current.next.next
    }
    current = current.next
  }
  return head
}

const head = {
  val: 1,
  next: {
    val: 1,
    next: {
      val: 2,
      next: {
        val: 3,
        next: {
          val: 3,
          next: null,
        },
      },
    },
  },
}
