import {test, expect} from "bun:test";
import {groupAnagrams} from "./group";

test("should group anagrams from a list of words", () => {
  let input = ["eat", "tea", "tan", "ate", "nat", "bat"];
  let output = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]];
  let got = groupAnagrams(input);
  console.log("got", got);
  expect(got).toEqual(output);
});

test("should handle single empty string as input", () => {
  let input = [""];
  let output = [[""]];
  expect(groupAnagrams(input)).toEqual(output);
});

test("should handle single character string as input", () => {
  let input = ["a"];
  let output = [["a"]];
  expect(groupAnagrams(input)).toEqual(output);
});

test("should group multiple empty strings together", () => {
  let input = ["", ""];
  let output = [["", ""]];
  expect(groupAnagrams(input)).toEqual(output);
});
