import {solution} from "../hamming_distance/solution_test.ts"
import {assertEquals} from "../../../dev_deps.ts"

Deno.test("difference of he whole word", () => {
  const a = "abcde"
  const b = "bcdef"
  const expected = 5
  const actual = solution(a, b)
  assertEquals(actual, expected)
})

Deno.test("difference of 1", () => {
  const a = "marcell"
  const b = "mercell"
  const expected = 1
  const actual = solution(a, b)
  assertEquals(actual, expected)
})
Deno.test("difference of 0", () => {
  const a = ""
  const b = ""
  const expected = 0
  const actual = solution(a, b)
  assertEquals(actual, expected)
})
Deno.test("difference of 0", () => {
  const a = "apa"
  const b = "apa"
  const expected = 0
  const actual = solution(a, b)
  assertEquals(actual, expected)
})
