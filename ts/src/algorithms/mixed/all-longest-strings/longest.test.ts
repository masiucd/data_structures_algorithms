import {test, expect} from "bun:test";
import {allLongestStrings, allLongestStringsTwo} from "./longest";

test("allLongestStrings to be ['aba', 'vcd', 'aba']", () => {
  let want = ["aba", "vcd", "aba"];
  let got = allLongestStrings(["aba", "aa", "ad", "vcd", "aba"]);
  expect(got).toStrictEqual(want);
});

test("allLongestStringsTwo to be ['aba', 'vcd', 'aba']", () => {
  let want = ["aba", "vcd", "aba"];
  let got = allLongestStringsTwo(["aba", "aa", "ad", "vcd", "aba"]);
  expect(got).toStrictEqual(want);
});
