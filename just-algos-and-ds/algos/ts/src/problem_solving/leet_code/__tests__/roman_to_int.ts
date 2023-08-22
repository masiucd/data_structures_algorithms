import {assertEquals} from "~/dev_deps.ts"
import {fromRomanToInt} from "../roman_to_int/solution.ts"

Deno.test("fromRomanToInt it works with fully list", () => {
  assertEquals(fromRomanToInt("LVIII"), 58)
})

Deno.test("fromRomanToInt it works with empty string", () => {
  assertEquals(fromRomanToInt("MCMXCIV"), 1994)
})

Deno.test("fromRomanToInt it works with empty string", () => {
  assertEquals(fromRomanToInt("III"), 3)
})
