import {test, expect} from "bun:test";
import {reverseVowels} from "./reverse";

test("reverseVowels", () => {
  let res = reverseVowels("hello");
  expect(res).toBe("holle");
});

test("reverseVowels", () => {
  let res = reverseVowels("leetcode");
  expect(res).toBe("leotcede");
});

test("reverseVowels", () => {
  let res = reverseVowels("aA");
  expect(res).toBe("Aa");
});
