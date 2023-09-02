import {isPalindrome, isPalindromeRecursive} from "./is_palindrome";
import {expect, test} from "bun:test";

test("isPalindrome should return true for abba", () => {
  expect(isPalindrome("abba")).toBe(true);
});

test("isPalindrome should return false for abbc", () => {
  expect(isPalindrome("abbc")).toBe(false);
});

test("isPalindromeRecursive should return true for abba", () => {
  expect(isPalindromeRecursive("abba")).toBe(true);
});

test("isPalindromeRecursive should return false for abbc", () => {
  expect(isPalindromeRecursive("abbc")).toBe(false);
});
