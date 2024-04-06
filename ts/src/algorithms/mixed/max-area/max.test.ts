import {test, expect} from "bun:test";

import {maxArea, maxAreaLinnear} from "./max";

test("Should return the max area", () => {
  expect(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])).toEqual(49);
  expect(maxAreaLinnear([1, 8, 6, 2, 5, 4, 8, 3, 7])).toEqual(49);
});

test("Should return the max area", () => {
  expect(maxArea([1, 1])).toEqual(1);
  expect(maxAreaLinnear([1, 1])).toEqual(1);
});
