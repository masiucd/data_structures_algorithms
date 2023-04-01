export interface ListNode {
  value: number
  next: ListNode | null
}

export function removeDuplicatesFromSortedList(head: ListNode) {
  let current: ListNode | null = head
  while (current) {
    while (current.next !== null && current.next.value === current.value) {
      current.next = current.next.next
    }
    current = current.next
  }
  return head
}

export function removeDuplicates(head: ListNode) {
  let current: ListNode | null = head
  while (current && current.next !== null) {
    if (current.value === current.next.value) {
      current.next = current.next.next
    }
    current = current.next
  }
  return head
}

export function getNodes(node: ListNode) {
  const xs: number[] = []
  let current: ListNode | null = node
  while (current !== null) {
    xs.push(current.value)
    current = current.next
  }
  return xs
}
