import {test, expect} from "bun:test";
import {hasPathSum} from "./sum";

test("hasPathSum should return true when there is a path that sums to target", () => {
  const tree = {
    val: 5,
    left: {
      val: 4,
      left: {
        val: 11,
        left: {val: 7, left: null, right: null},
        right: {val: 2, left: null, right: null},
      },
      right: null,
    },
    right: {
      val: 8,
      left: {val: 13, left: null, right: null},
      right: {
        val: 4,
        left: null,
        right: {val: 1, left: null, right: null},
      },
    },
  };
  expect(hasPathSum(tree, 22)).toBe(true);
});

test("hasPathSum should return false when there is no path that sums to target", () => {
  const tree = {
    val: 1,
    left: {
      val: 2,
      left: null,
      right: null,
    },
    right: null,
  };
  expect(hasPathSum(tree, 1)).toBe(false);
});
