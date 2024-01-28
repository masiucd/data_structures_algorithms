import {test, expect} from "bun:test";
import {alternatingSums, alternatingSumsV2} from "./sums";

test("should return correct alternating sums for array with 5 elements", () => {
  const actual = alternatingSums([50, 60, 60, 45, 70]);
  expect(actual).toEqual([180, 105]);
});

test("should return correct alternating sums for array with 4 elements", () => {
  const actual = alternatingSumsV2([100, 50, 50, 100]);
  expect(actual).toEqual([150, 150]);
});
