import {assertEquals} from "~/dev_deps.ts"
import {
  containsNearbyDuplicate,
  containsNearbyDuplicate2,
} from "../contains_near_by_duplicate/solution.ts"

Deno.test("containsNearbyDuplicate it works", () => {
  assertEquals(containsNearbyDuplicate([1, 2, 3, 1], 3), true)
  assertEquals(containsNearbyDuplicate2([1, 2, 3, 1], 3), true)
})

Deno.test("containsNearbyDuplicate it works", () => {
  assertEquals(containsNearbyDuplicate([1, 0, 1, 1], 1), true)
  assertEquals(containsNearbyDuplicate2([1, 0, 1, 1], 1), true)
})

Deno.test("containsNearbyDuplicate it works", () => {
  assertEquals(containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2), false)
  assertEquals(containsNearbyDuplicate2([1, 2, 3, 1, 2, 3], 2), false)
})
