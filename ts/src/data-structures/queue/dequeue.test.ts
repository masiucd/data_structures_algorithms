import {test, expect} from "bun:test";
import {Deque} from "./dequeue";

test("Deque", () => {
  let deque = new Deque();
  expect(deque.isEmpty()).toBe(true);
  deque.append(1);
  expect(deque.isEmpty()).toBe(false);
  deque.appendleft(2);
  expect(deque.pop()).toBe(1);
  expect(deque.popleft()).toBe(2);
  expect(deque.popleft()).toBe(-1);
});
