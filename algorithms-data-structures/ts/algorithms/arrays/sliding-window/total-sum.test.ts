import {test, expect} from "bun:test";
import {totalSum} from "./total-sum";

test("totalSum", () => {
  expect(totalSum([1, 2, 3, 4], 2)).toBe(7);
  expect(totalSum([10, 1, 2, 3, 4], 3)).toBe(13);
});
