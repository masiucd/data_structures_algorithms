import {test, expect} from "bun:test";

import {sortByHeight} from "./sort";

test("should sort heights", () => {
  let input = [-1, 150, 190, 170, -1, -1, 160, 180];
  let output = [-1, 150, 160, 170, -1, -1, 180, 190];
  expect(sortByHeight(input)).toEqual(output);
});
