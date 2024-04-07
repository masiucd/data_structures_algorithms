import {test, expect} from "bun:test";
import {missingNumber} from "./missing";

test("missingNumber", () => {
  let nums = [3, 0, 1];
  let res = missingNumber(nums);
  expect(res).toBe(2);

  nums = [0, 1];
  res = missingNumber(nums);
  expect(res).toBe(2);

  nums = [9, 6, 4, 2, 3, 5, 7, 0, 1];
  res = missingNumber(nums);
  expect(res).toBe(8);
});
