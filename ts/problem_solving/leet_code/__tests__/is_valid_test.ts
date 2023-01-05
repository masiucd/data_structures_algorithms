import {assertEquals} from "../../../dev_deps.ts"
import {brackets, isValid} from "../is_valid/is_valid.ts"

Deno.test("isValid it works with fully list", () => {
  assertEquals(isValid("()"), true)
  assertEquals(isValid("()[]{}"), true)
  assertEquals(isValid("(]"), false)
})

Deno.test(
  "brackets store close brackets as keys and open brackets as values",
  () => {
    assertEquals(brackets["}"], "{")
    assertEquals(brackets["]"], "[")
    assertEquals(brackets[")"], "(")
  }
)
