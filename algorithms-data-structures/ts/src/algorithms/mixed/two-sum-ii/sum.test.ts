import {test, expect} from "bun:test";
import {twoSum} from "./sum";

test("Should return indices of two numbers that add up to 9", () => {
  expect(twoSum([2, 7, 12, 33], 9)).toEqual([1, 2]);
});

test("Should return indices of two numbers that add up to 6", () => {
  expect(twoSum([2, 3, 4], 6)).toEqual([1, 3]);
});

test("Should return indices of two numbers that add up to -1", () => {
  expect(twoSum([-1, 0], -1)).toEqual([1, 2]);
});
