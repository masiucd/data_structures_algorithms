import {assertEquals} from "~/dev_deps.ts"
import {
  mergeSortedList,
  ListNode,
  printNodes,
} from "../merge_sorted_list/solution.ts"

Deno.test("mergeSortedList it works", () => {
  assertEquals(
    mergeSortedList(
      new ListNode(1, new ListNode(2, new ListNode(4))),
      new ListNode(1, new ListNode(3, new ListNode(4)))
    ),
    new ListNode(
      1,
      new ListNode(
        1,
        new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(4))))
      )
    )
  )
})
Deno.test("mergeSortedList it works", () => {
  const res = mergeSortedList(
    new ListNode(1, new ListNode(2, new ListNode(4))),
    new ListNode(1, new ListNode(3, new ListNode(4)))
  )
  const got = printNodes(res as ListNode)
  const want = [1, 1, 2, 3, 4, 4]
  assertEquals(got, want)
})

// console.log("res", printNodes(res as ListNode))
