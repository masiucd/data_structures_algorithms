import {commonCharacterCount} from "./common";
import {test, expect} from "bun:test";

test("commonCharacterCount", () => {
  let s1 = "aabcc";
  let s2 = "adcaa";
  let res = commonCharacterCount(s1, s2); // 3
  expect(res).toEqual(3);
});

test("commonCharacterCount", () => {
  let s1 = "zzzz";
  let s2 = "zzzzzzz";
  let res = commonCharacterCount(s1, s2); // 4
  expect(res).toEqual(4);
});
