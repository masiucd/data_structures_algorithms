import {test, expect} from "bun:test";
import {adjacentElementsProduct} from "./adjacent";

test("adjacentElementsProduct", () => {
  expect(adjacentElementsProduct([3, 6, -2, -5, 7, 3])).toBe(21);
  expect(adjacentElementsProduct([9, 5, 10, 2, 24, -1, -48])).toBe(50);
  expect(adjacentElementsProduct([5, 1, 2, 3, 1, 4])).toBe(6);
  expect(adjacentElementsProduct([1, 2, 3, 0])).toBe(6);
  expect(adjacentElementsProduct([1, 0, 1, 0, 1000])).toBe(0);
});
