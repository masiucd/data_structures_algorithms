import {expect, test} from "bun:test";
import {Solution} from "./reverse";
import {ListNode} from "../list-node";

test("Reverse Linked List", () => {
  let solution = new Solution();
  let head = new ListNode(1);
  head.next = new ListNode(2);
  head.next.next = new ListNode(3);
  let result = solution.reverseList(head);
  expect(result?.val).toBe(3);
  expect(result?.next?.val).toBe(2);
  expect(result?.next?.next?.val).toBe(1);
});
