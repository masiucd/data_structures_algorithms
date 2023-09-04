import {sum, sumRecursive, sumReduce} from "./sum";
import {expect, test} from "bun:test";

test("sum the list [1,2,3,4,5] should be 15", () => {
  expect(sumRecursive([1, 2, 3, 4, 5])).toBe(15);
  expect(sum([1, 2, 3, 4, 5])).toBe(15);
  expect(sumReduce([1, 2, 3, 4, 5])).toBe(15);
});
