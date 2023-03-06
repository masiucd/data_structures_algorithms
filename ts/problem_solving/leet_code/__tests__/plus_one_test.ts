import {assertEquals} from "../../../dev_deps.ts"
import {plusOne, plusOne2} from "../plus_one/solution.ts"

Deno.test("plusOne it works with fully list", () => {
  assertEquals(plusOne([1, 2, 3]), [1, 2, 4])
})

Deno.test("plusOne it works with empty string", () => {
  assertEquals(plusOne([9]), [1, 0])
})

Deno.test("plusOne2 it works with fully list", () => {
  assertEquals(plusOne2([1, 2, 3]), [1, 2, 4])
})

Deno.test("plusOne2 it works with empty string", () => {
  assertEquals(plusOne2([9]), [1, 0])
})
