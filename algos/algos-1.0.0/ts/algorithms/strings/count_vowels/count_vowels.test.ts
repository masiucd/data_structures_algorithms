import {expect, test} from "bun:test";
import {
  countVowels,
  countVowelsFunctional,
  countVowelsIterative,
} from "./count_vowels";

test("countVowels should return 0 for empty string", () => {
  expect(countVowels("")).toBe(0);

  expect(countVowelsIterative("")).toBe(0);

  expect(countVowelsFunctional("")).toBe(0);
});

test("countVowels should return 1 for a", () => {
  expect(countVowels("a")).toBe(1);

  expect(countVowelsIterative("a")).toBe(1);

  expect(countVowelsFunctional("a")).toBe(1);
});

test("countVowels should return 0 for b", () => {
  expect(countVowels("b")).toBe(0);

  expect(countVowelsIterative("b")).toBe(0);

  expect(countVowelsFunctional("b")).toBe(0);
});

test("countVowels should return 1 for ab", () => {
  expect(countVowels("ab")).toBe(1);

  expect(countVowelsIterative("ab")).toBe(1);

  expect(countVowelsFunctional("ab")).toBe(1);
});

test("countVowels should return 2 for aa", () => {
  expect(countVowels("aa")).toBe(2);

  expect(countVowelsIterative("aa")).toBe(2);

  expect(countVowelsFunctional("aa")).toBe(2);
});

test("countVowels should return 2 for marcell", () => {
  expect(countVowels("marcell")).toBe(2);

  expect(countVowelsIterative("marcell")).toBe(2);

  expect(countVowelsFunctional("marcell")).toBe(2);
});
