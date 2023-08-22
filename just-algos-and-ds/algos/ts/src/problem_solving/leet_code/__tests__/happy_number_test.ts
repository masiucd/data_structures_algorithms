import {assertEquals, assert, assertFalse} from "~/dev_deps.ts"
import {isHappy, isHappy2, numOfSquare} from "../happy_number/solution.ts"

Deno.test("isHappy it works 19", () => {
  assertEquals(isHappy(19), true)
  assert(isHappy2(19))
})

Deno.test("isHappy it works 2", () => {
  assertEquals(isHappy(2), false)
  assertFalse(isHappy2(2))
})

Deno.test("isHappy it works 1", () => {
  assertEquals(isHappy(1), true)
  assert(isHappy2(1))
})

Deno.test("isHappy it works 14", () => {
  assertEquals(isHappy(14), false)
  assertFalse(isHappy2(14))
})

Deno.test("numOfSquare it works 19", () => {
  assertEquals(numOfSquare(19), 82)
})

Deno.test("numOfSquare it works 2", () => {
  assertEquals(numOfSquare(2), 4)
})

Deno.test("numOfSquare it works 82", () => {
  assertEquals(numOfSquare(82), 68)
})
