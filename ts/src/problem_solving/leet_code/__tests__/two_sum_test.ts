import {assertEquals} from "../../../dev_deps.ts"
import {twoSum} from "../two_sum/solution.ts"

Deno.test("twoSum it works", () => {
  assertEquals(twoSum([2, 7, 11, 15], 9), [0, 1])
  assertEquals(twoSum([3, 2, 4], 6), [1, 2])
  assertEquals(twoSum([3, 3], 6), [0, 1])
  assertEquals(twoSum([1, 2, 3], 12), [0, 0])
})
