import {assertEquals} from "~/dev_deps.ts"
import {brackets, isValid, isValidTwo} from "../is_valid/is_valid.ts"

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

Deno.test("isValidTwo it works with fully list", () => {
  assertEquals(isValidTwo("()"), true)
  assertEquals(isValidTwo("()[]{}"), true)
  assertEquals(isValidTwo("(]"), false)
})

Deno.test("isValidTwo it works with partially list", () => {
  assertEquals(isValidTwo("()"), true)
  assertEquals(isValidTwo("()[]{}"), true)
  assertEquals(isValidTwo("(]"), false)
  assertEquals(isValidTwo("([)]"), false)
  assertEquals(isValidTwo("{[]}"), true)
})
