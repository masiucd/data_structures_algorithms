import {test, expect} from "bun:test";
import {reverseInParentheses} from "./reverse";

test("should reverse characters in parentheses", () => {
  let actual = reverseInParentheses("foo(bar)baz");
  expect(actual).toBe("foorabbaz");
});

test("should reverse characters in multiple parentheses", () => {
  let actual = reverseInParentheses("foo(bar)baz(blim)");
  expect(actual).toBe("foorabbazmilb");
});

test("should reverse characters in nested parentheses", () => {
  let actual = reverseInParentheses("foo(bar(baz))blim");
  expect(actual).toBe("foobazrabblim");
});

test("should return empty string when input is empty", () => {
  let actual = reverseInParentheses("");
  expect(actual).toBe("");
});

test("should return empty string when parentheses are empty", () => {
  let actual = reverseInParentheses("()");
  expect(actual).toBe("");
});

test("should reverse characters in parentheses and ignore characters outside", () => {
  let actual = reverseInParentheses("(abc)d(efg)");
  expect(actual).toBe("cbadgfe");
});
