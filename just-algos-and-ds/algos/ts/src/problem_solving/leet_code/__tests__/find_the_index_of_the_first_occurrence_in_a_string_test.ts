import {assertEquals} from "~/dev_deps.ts"
import {solution} from "../find_the_index_of_the_first_occurrence_in_a_string/solution.ts"

Deno.test("solution it works", () => {
  assertEquals(solution("hello", "ll"), 2)
})

Deno.test("solution it works", () => {
  assertEquals(solution("aaaaa", "bba"), -1)
})
