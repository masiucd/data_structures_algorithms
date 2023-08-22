import {assertEquals} from "~/dev_deps.ts"
import {
  solution,
  solutionTwo,
  solutionThree,
} from "../number_of_bits/solution.ts"

Deno.test("number of bits solution one", () => {
  assertEquals(solution(0b00000000000000000000000000001011), 3)
  assertEquals(solution(0b00000000000000000000000010000000), 1)
  assertEquals(solution(0b11111111111111111111111111111101), 31)
})

Deno.test("number of bits solution two", () => {
  assertEquals(solutionTwo(0b00000000000000000000000000001011), 3)
  assertEquals(solutionTwo(0b00000000000000000000000010000000), 1)
  assertEquals(solutionTwo(0b11111111111111111111111111111101), 31)
})

Deno.test("number of bits solution three", () => {
  assertEquals(solutionThree(0b00000000000000000000000000001011), 3)
  assertEquals(solutionThree(0b00000000000000000000000010000000), 1)
  assertEquals(solutionThree(0b11111111111111111111111111111101), 31)
})
