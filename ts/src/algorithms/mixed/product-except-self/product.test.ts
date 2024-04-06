import {test, expect} from "bun:test";

import {productExceptSelf} from "./product";

test("productExceptSelf", () => {
  let input = [1, 2, 3, 4];
  let output = [24, 12, 8, 6];
  expect(productExceptSelf(input)).toEqual(output);
});

test("productExceptSelf", () => {
  let input = [-1, 1, 0, -3, 3];
  let output = [0, 0, 9, 0, 0];
  expect(productExceptSelf(input)).toEqual(output);
  // expect(productExceptSelfTwo(input)).toEqual(output);
});
