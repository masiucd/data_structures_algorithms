import {test, expect} from "bun:test";
import {matrixElementsSum} from "./sum";

test("matrixElementsSum to be 9", () => {
  expect(
    matrixElementsSum([
      [0, 1, 1, 2],
      [0, 5, 0, 0],
      [2, 0, 3, 3],
    ])
  ).toBe(9);
});
