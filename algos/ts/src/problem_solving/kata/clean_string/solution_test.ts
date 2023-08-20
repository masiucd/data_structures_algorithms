import {assertEquals} from "../../../dev_deps.ts"
import {solution} from "./solution.ts"

Deno.test("solution it works", () => {
  assertEquals(solution("abc#d##c"), "ac")
  assertEquals(solution("abc##d######"), "")
  assertEquals(solution("#######"), "")
})
