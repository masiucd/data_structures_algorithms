import {MinStack} from "./min-stack";
import {expect, test} from "bun:test";

test("Minstack", () => {
  let minStack = new MinStack();
  minStack.push(1);
  minStack.push(-2);
  minStack.push(3);
  expect(minStack.getMin()).toBe(-2);
  minStack.pop();
  expect(minStack.top()).toBe(-2);
  expect(minStack.getMin()).toBe(-2);
  minStack.pop();
  expect(minStack.getMin()).toBe(1);
});
