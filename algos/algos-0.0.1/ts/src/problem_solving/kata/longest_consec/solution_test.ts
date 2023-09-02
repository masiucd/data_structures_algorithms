import {assertEquals} from "../../../dev_deps.ts"
import {solution} from "./solution.ts"

Deno.test(function longestConsentTestItWorks() {
  assertEquals(
    solution(["zone", "abigail", "theta", "form", "libe", "zas"], 2),
    "abigailtheta"
  )
  assertEquals(
    solution(["zone", "abigail", "theta", "form", "libe", "zas"], 3),
    "zoneabigailtheta"
  )
})
