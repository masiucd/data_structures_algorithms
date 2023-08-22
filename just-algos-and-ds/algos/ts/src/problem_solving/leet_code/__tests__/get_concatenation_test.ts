import {getConcatenation} from "../get_concatenation/solution.ts"
import {assertEquals} from "~/dev_deps.ts"

Deno.test("getConcatenation it works with fully list", () => {
  assertEquals(
    getConcatenation([1, 2, 3, 4, 5]),
    [1, 2, 3, 4, 5, 1, 2, 3, 4, 5]
  )
})

Deno.test("getConcatenation it works with empty string", () => {
  assertEquals(getConcatenation([]), [])
})
