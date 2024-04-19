import {expect, test} from "bun:test";
import {findContentChildren} from "./find";

test("g = [1, 2], s = [1, 2, 3] should be 2", () => {
  expect(findContentChildren([1, 2], [1, 2, 3])).toBe(2);
});

test("g = [1, 2, 3], s = [1, 1] should be 1", () => {
  expect(findContentChildren([1, 2, 3], [1, 1])).toBe(1);
});
