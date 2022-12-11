import {assertEquals} from "../../../dev_deps.ts"
import {exponent} from "./exponent.ts"

Deno.test(function exponentTestItWorks() {
  assertEquals(exponent(2, 0), 1)
  assertEquals(exponent(2, 3), 8)
  assertEquals(exponent(2, 4), 16)
  assertEquals(exponent(5, 2), 25)
})
