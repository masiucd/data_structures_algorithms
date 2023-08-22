import {assertEquals} from "~/dev_deps.ts"

import {isSameTree} from "../is_same_tree/solution.ts"

Deno.test("isSameTree works", () => {
  const p = {
    val: 1,
    left: {
      val: 2,
      left: null,
      right: null,
    },
    right: {
      val: 3,
      left: null,
      right: null,
    },
  }
  const q = {
    val: 1,
    left: {
      val: 2,
      left: null,
      right: null,
    },
    right: {
      val: 3,
      left: null,
      right: null,
    },
  }
  assertEquals(isSameTree(p, q), true)
})
