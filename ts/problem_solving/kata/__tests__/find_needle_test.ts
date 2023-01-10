import {assertEquals} from "../../../dev_deps.ts"
import {solution} from "../find_needle/solution.ts"

Deno.test("findNeedle it works when there is needle!", () => {
  const haystack = ["hay", "needle", "hay", "hay", "more hay", "randomJunk"]
  const result = solution(haystack)
  assertEquals(result, "found the needle at position 1")
})

Deno.test("findNeedle it works when there is NO needle!", () => {
  const haystack = ["hay", "banana", "hay", "hay", "more hay", "randomJunk"]
  const result = solution(haystack)
  assertEquals(result, "")
})
