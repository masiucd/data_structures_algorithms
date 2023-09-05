import {expect, test} from "bun:test";
import {removeElement, removeElementTwo} from "./remove_element";

test("removeElement", () => {
  expect(removeElement([3, 2, 2, 3], 3)).toBe(2);
  expect(removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2)).toBe(5);
});

test("removeElementTwo", () => {
  expect(removeElementTwo([3, 2, 2, 3], 3)).toBe(2);
  expect(removeElementTwo([0, 1, 2, 2, 3, 0, 4, 2], 2)).toBe(5);
});
