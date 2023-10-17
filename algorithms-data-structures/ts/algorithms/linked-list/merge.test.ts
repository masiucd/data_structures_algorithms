import {merge, printList, Node} from "./merge";
import {test, expect} from "bun:test";

test("merge", () => {
  let t1 = new Node(1);
  t1.next = new Node(3);
  t1.next.next = new Node(5);

  let t2 = new Node(2);
  t2.next = new Node(4);
  t2.next.next = new Node(6);

  let merged = merge(t1, t2);

  expect(printList(merged)).toEqual([1, 2, 3, 4, 5, 6]);
});
