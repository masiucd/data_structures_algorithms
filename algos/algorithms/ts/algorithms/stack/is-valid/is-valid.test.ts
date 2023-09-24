import {expect, test} from "bun:test";
import {isValid} from "./is-valid";

test("() should be true", () => {
  expect(isValid("()")).toBe(true);
});

test("()[]{} should be true", () => {
  expect(isValid("()[]{}")).toBe(true);
});

test("(] should be false", () => {
  expect(isValid("(]")).toBe(false);
});

test("([)] should be false", () => {
  expect(isValid("([)]")).toBe(false);
});

test("{[]} should be true", () => {
  expect(isValid("{[]}")).toBe(true);
});

test("] should be false", () => {
  expect(isValid("]")).toBe(false);
});

test("(( should be false", () => {
  expect(isValid("((")).toBe(false);
});

test("((())) should be true", () => {
  expect(isValid("((()))")).toBe(true);
});
