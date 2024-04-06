import {test, expect} from "bun:test";
import {reverseList, reverseListTwo} from "./reverse";

test("reverseList", () => {
  expect(reverseList([1, 2, 3, 4, 5])).toEqual([5, 4, 3, 2, 1]);
  expect(reverseListTwo([1, 2, 3, 4, 5])).toEqual([5, 4, 3, 2, 1]);
});
