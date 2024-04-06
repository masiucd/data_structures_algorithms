import {test, expect} from "bun:test";

import {boxBlur} from "./box-blur";

test("boxBlur", () => {
  expect(
    boxBlur([
      [1, 1, 1],
      [1, 7, 1],
      [1, 1, 1],
    ])
  ).toEqual([[1]]);
  expect(
    boxBlur([
      [7, 4, 0, 1],
      [5, 6, 2, 2],
      [6, 10, 7, 8],
      [1, 4, 2, 0],
    ])
  ).toEqual([
    [5, 4],
    [4, 4],
  ]);
});
