import {test, expect} from "bun:test";
import {Stack} from "./stack";

test("Stack", () => {
  let stack = new Stack();
  expect(stack.isEmpty()).toBe(true);
  stack.push(5);
  stack.push(8);
  expect(stack.peek()).toBe(8);
  stack.push(11);
  expect(stack.size()).toBe(3);
  expect(stack.isEmpty()).toBe(false);
  stack.push(15);
  stack.pop();
  stack.pop();
  expect(stack.size()).toBe(2);
  expect(stack.peek()).toBe(8);
  stack.pop();
  stack.pop();
  expect(stack.isEmpty()).toBe(true);
  expect(stack.peek()).toBe(undefined);
  stack.push(1);
  expect(stack.peek()).toBe(1);
});
