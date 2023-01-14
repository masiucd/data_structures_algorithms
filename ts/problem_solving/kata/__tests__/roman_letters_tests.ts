import {assertEquals} from "../../../dev_deps.ts"
import {solution} from "../roman_letters/solution.ts"

Deno.test("romanLetters it works when there is IV!", () => {
  const roman = "IV"
  const result = solution(roman)
  assertEquals(result, 4)
})

Deno.test("romanLetters it works when there is XXI!", () => {
  const roman = "XXI"
  const result = solution(roman)
  assertEquals(result, 21)
})

Deno.test("romanLetters it works when there is I!", () => {
  const roman = "I"
  const result = solution(roman)
  assertEquals(result, 1)
})

Deno.test("romanLetters it works when there is MMVIII!", () => {
  const roman = "MMVIII"
  const result = solution(roman)
  assertEquals(result, 2008)
})

Deno.test("romanLetters it works when there is MDCLXVI!", () => {
  const roman = "MDCLXVI"
  const result = solution(roman)
  assertEquals(result, 1666)
})
