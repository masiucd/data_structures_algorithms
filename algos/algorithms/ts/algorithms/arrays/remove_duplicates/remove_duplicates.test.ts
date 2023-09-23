import {test, expect} from "bun:test";
import {removeDuplicates, removeDuplicatesRec} from "./remove_duplicates";

test("remove duplicates from [1,1,2] should be 2", () => {
  expect(removeDuplicates([1, 1, 2])).toBe(2);
  expect(removeDuplicatesRec([1, 1, 2], new Set())).toBe(2);
});

test("remove duplicates from [0,0,1,1,1,2,2,3,3,4] should be 5", () => {
  expect(removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])).toBe(5);
  expect(removeDuplicatesRec([0, 0, 1, 1, 1, 2, 2, 3, 3, 4], new Set())).toBe(
    5
  );
});
