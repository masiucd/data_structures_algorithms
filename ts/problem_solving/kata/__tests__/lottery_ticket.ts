import {assertEquals} from "../../../dev_deps.ts"
import {bingo} from "../lottery_ticket/solution.ts"

Deno.test("bingo it works!", () => {
  const res = bingo(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    2
  )
  assertEquals(res, "Loser!")
})

Deno.test("bingo it works!", () => {
  const res = bingo(
    [
      ["ABC", 65],
      ["HGR", 74],
      ["BYHT", 74],
    ],
    1
  )
  assertEquals(res, "Winner!")
})
