import {test, expect} from "bun:test";
import {insertionSort, insertionSortUsingSwapHelper} from "./sort";

test("insertionSort", () => {
  expect(insertionSort([1, 2, 3, 4])).toEqual([1, 2, 3, 4]);
  expect(insertionSortUsingSwapHelper([1, 2, 3, 4])).toEqual([1, 2, 3, 4]);

  expect(insertionSort([10, 1, 2, 3, 4])).toEqual([1, 2, 3, 4, 10]);
  expect(insertionSortUsingSwapHelper([10, 1, 2, 3, 4])).toEqual([
    1, 2, 3, 4, 10,
  ]);

  expect(insertionSort([5, 4, 3, 2, 1])).toEqual([1, 2, 3, 4, 5]);
  expect(insertionSortUsingSwapHelper([5, 4, 3, 2, 1])).toEqual([
    1, 2, 3, 4, 5,
  ]);

  expect(insertionSort([])).toEqual([]);
  expect(insertionSortUsingSwapHelper([])).toEqual([]);
});
