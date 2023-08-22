/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

export class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val
    this.next = next === undefined ? null : next
  }
}

export function mergeSortedList(
  list1: ListNode | null,
  list2: ListNode | null
): ListNode | null {
  if (list1 === null) return list2
  const newHead = new ListNode()
  let tail = newHead
  while (list1 && list2) {
    if (list1.val < list2.val) {
      tail.next = list1
      list1 = list1.next
    } else {
      tail.next = list2
      list2 = list2.next
    }
    tail = tail.next
  }
  if (list1) {
    tail.next = list1
  } else {
    tail.next = list2
  }
  return newHead.next
}

export function printNodes(node: ListNode) {
  let current: ListNode | null = node
  const xs = []
  while (current !== null) {
    xs.push(current.val)
    current = current.next
  }
  return xs
}
