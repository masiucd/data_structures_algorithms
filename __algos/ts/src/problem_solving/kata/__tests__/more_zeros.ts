import {assertEquals} from "../../../dev_deps.ts"
import {
  moreZeros,
  charToBinary,
  countZerosAndOnesInBinary,
} from "../more_zeros/solution.ts"

Deno.test("moreZeros it works", () => {
  let actual = moreZeros("abcdeea")
  let expected = ["a", "b", "d"]
  assertEquals(actual, expected)

  actual = moreZeros("Great job!")
  expected = ["a", " ", "b", "!"]
  assertEquals(actual, expected)
})

Deno.test("charToBinary", () => {
  const actual = charToBinary("a")
  const expected = "1100001"
  assertEquals(actual, expected)
})

Deno.test("countZerosAndOnesInBinary", () => {
  const actual = countZerosAndOnesInBinary("1100001")
  const expected = [4, 3]
  assertEquals(actual, expected)
})
