import {assertEquals} from "~/dev_deps.ts"

import {
  removeDuplicatesFromSortedList,
  removeDuplicates,
  getNodes,
} from "../remove_duplicates_sorted_list/solution.ts"

const head = {
  value: 1,
  next: {
    value: 1,
    next: {
      value: 2,
      next: {
        value: 3,
        next: {
          value: 3,
          next: null,
        },
      },
    },
  },
}

Deno.test("removeDuplicatesFromSortedList it works", () => {
  assertEquals(removeDuplicatesFromSortedList(head), {
    value: 1,
    next: {
      value: 2,
      next: {
        value: 3,
        next: null,
      },
    },
  })
})

Deno.test("removeDuplicates it works", () => {
  assertEquals(removeDuplicates(head), {
    value: 1,
    next: {
      value: 2,
      next: {
        value: 3,
        next: null,
      },
    },
  })
})

Deno.test("getNodes it works", () => {
  assertEquals(getNodes(head), [1, 1, 2, 3, 3])
})
