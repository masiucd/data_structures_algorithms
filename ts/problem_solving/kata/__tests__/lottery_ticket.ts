import {assertEquals} from "../../../dev_deps.ts"
import {solution, solution2} from "../lottery_ticket/solution.ts"

Deno.test("solution1 it works!", () => {
  const res = solution(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    2
  )
  assertEquals(res, "Loser!")
})

Deno.test("solution2 it works!", () => {
  const res = solution(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    1
  )
  assertEquals(res, "Winner!")
})

Deno.test("solution2 it works!", () => {
  const res = solution2(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    2
  )
  assertEquals(res, "Loser!")
})

Deno.test("solution2 it works!", () => {
  const res = solution2(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    1
  )
  assertEquals(res, "Winner!")
})
