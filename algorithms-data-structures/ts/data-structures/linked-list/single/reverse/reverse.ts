import {ListNode} from "../list-node";

export class Solution {
  reverseList(head: ListNode | null): ListNode | null {
    let prev = null;
    let curr = head;
    while (curr !== null) {
      let nextTemp = curr.next;
      curr.next = prev;
      prev = curr;
      curr = nextTemp;
    }
    return prev;
  }
}
