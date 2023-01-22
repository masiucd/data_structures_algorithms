import {assertEquals} from "../../../dev_deps.ts"
import {flattenDeep} from "../flat_list/solution.ts"

Deno.test("flattenDeep it works with fully list", () => {
  assertEquals(flattenDeep([1, [2, [3, [4]], 5]]), [1, 2, 3, 4, 5])
})

Deno.test("flattenDeep with a NON nested list", () => {
  assertEquals(flattenDeep([1, 2, 3]), [1, 2, 3])
})
