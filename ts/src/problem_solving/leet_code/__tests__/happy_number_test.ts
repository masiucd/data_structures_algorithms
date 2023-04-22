import {assertEquals} from "~/dev_deps.ts"
import {isHappy} from "../happy_number/solution.ts"

Deno.test("isHappy it works 19", () => {
  assertEquals(isHappy(19), true)
})

Deno.test("isHappy it works 2", () => {
  assertEquals(isHappy(2), false)
})

Deno.test("isHappy it works 1", () => {
  assertEquals(isHappy(1), true)
})

Deno.test("isHappy it works 14", () => {
  assertEquals(isHappy(14), false)
})
